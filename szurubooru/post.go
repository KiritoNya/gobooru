package szurubooru

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	Version            int           `json:"version,omitempty"`
	Id                 int           `json:"id,omitempty"`
	CreationTime       time.Time     `json:"creationTime,omitempty"`
	LastEditTime       time.Time     `json:"lastEditTime,omitempty"`
	Safety             string        `json:"safety"`
	Source             string        `json:"source,omitempty"`
	Type               string        `json:"type,omitempty"`
	Checksum           string        `json:"checksum,omitempty"`
	ChecksumMD5        string        `json:"checksumMD5,omitempty"`
	CanvasWidth        int           `json:"canvasWidth,omitempty"`
	CanvasHeight       int           `json:"canvasHeight,omitempty"`
	ContentUrl         string        `json:"contentUrl,omitempty"`
	ThumbnailUrl       string        `json:"thumbnailUrl,omitempty"`
	Flags              []string      `json:"flags,omitempty"`
	Tags               []*MicroTag   `json:"tags"`
	Relations          []*Post       `json:"relations,omitempty"`
	Notes              interface{}   `json:"notes,omitempty"`
	User               interface{}   `json:"user,omitempty"`
	Score              int           `json:"score,omitempty"`
	OwnScore           int           `json:"ownScore,omitempty"`
	OwnFavorite        bool          `json:"ownFavorite,omitempty"`
	TagCount           int           `json:"tagCount,omitempty"`
	FavoriteCount      int           `json:"favoriteCount,omitempty"`
	CommentCount       int           `json:"commentCount,omitempty"`
	NoteCount          int           `json:"noteCount,omitempty"`
	FeatureCount       int           `json:"featureCount,omitempty"`
	RelationCount      int           `json:"relationCount,omitempty"`
	LastFeatureTime    time.Time     `json:"lastFeatureTime,omitempty"`
	FavoritedBy        interface{}   `json:"favoritedBy,omitempty"`
	HasCustomThumbnail bool          `json:"hasCustomThumbnail,omitempty"`
	MimeType           string        `json:"mimeType,omitempty"`
	Comments           []interface{} `json:"comments,omitempty"`
	Pools              []interface{} `json:"pools,omitempty"`
}

type PostOperation struct {
	Post
	Tags      []string `json:"tags,omitempty"`
	Relations []int    `json:"relations,omitempty"`
}

type MicroPost struct {
	Name         string `json:"name"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

type Note struct {
	Polygon [][]int `json:"polygon"`
	Text    string  `json:"text"`
}

//Create is a method of Post object that create a post
func (p *Post) Create(fileData []byte, fileName string) (*Post, error) {
	// Check input validity
	if p.Safety == "" {
		return nil, errors.New("Safety must be setted")
	}

	// Upload file
	token, err := Upload(fileData, fileName)
	if err != nil {
		return nil, err
	}

	// Get tags name
	var listTags []string
	for _, tag := range p.Tags {
		listTags = append(listTags, tag.Names[0])
	}

	// Create Json
	var raw = struct {
		Safety       string   `json:"safety"`
		Tags         []string `json:"tags"`
		ContentToken string   `json:"contentToken"`
	}{
		Safety:       p.Safety,
		Tags:         listTags,
		ContentToken: token,
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}

	// Do request
	content, code, err := Client.Do("POST", "/posts", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Check status code
	if code != http.StatusOK {
		return nil, parseError(content)
	}

	// Parse json response
	var post Post
	err = json.Unmarshal(content, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// Update is a method of Post object that modify a post
func (p *Post) Update() (*Post, error) {
	if p.Id == 0 {
		return nil, errors.New("post id not found")
	}

	data, err := json.Marshal(p.toPostOperation())
	if err != nil {
		return nil, err
	}

	data, code, err := Client.Do("PUT", "/post/"+strconv.Itoa(p.Id), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Check status code
	if code != http.StatusOK {
		fmt.Println(code)
		return nil, parseError(data)
	}

	// Parse json response
	var post Post
	err = json.Unmarshal(data, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *Post) Delete() error {
	// Check params
	if p.Id == 0 {
		return errors.New("post id not found")
	}

	if p.Version == 0 {
		return errors.New("post version not setted")
	}

	// Create json request body
	raw := struct {
		Version int `json:"version"`
	}{
		p.Version,
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return err
	}

	// Do request
	data, code, err := Client.Do("PUT", "/post/"+strconv.Itoa(p.Id), bytes.NewReader(data))
	if err != nil {
		return err
	}

	// Check status code
	if code != http.StatusOK {
		return parseError(data)
	}

	return nil
}

func (p *Post) toPostOperation() *PostOperation {
	var tags []string
	var relIds []int

	for _, rel := range p.Relations {
		relIds = append(relIds, rel.Id)
	}

	for _, tag := range p.Tags {
		tags = append(tags, tag.Names[0])
	}

	return &PostOperation{
		Post: Post{
			Version:            p.Version,
			Id:                 p.Id,
			CreationTime:       p.CreationTime,
			LastEditTime:       p.LastEditTime,
			Safety:             p.Safety,
			Source:             p.Source,
			Type:               p.Type,
			Checksum:           p.Checksum,
			ChecksumMD5:        p.ChecksumMD5,
			CanvasWidth:        p.CanvasWidth,
			CanvasHeight:       p.CanvasHeight,
			ContentUrl:         p.ContentUrl,
			ThumbnailUrl:       p.ThumbnailUrl,
			Flags:              p.Flags,
			Notes:              p.Notes,
			User:               p.User,
			Score:              p.Score,
			OwnScore:           p.OwnScore,
			OwnFavorite:        p.OwnFavorite,
			TagCount:           p.TagCount,
			FavoriteCount:      p.FavoriteCount,
			CommentCount:       p.CommentCount,
			NoteCount:          p.NoteCount,
			FeatureCount:       p.FeatureCount,
			RelationCount:      p.RelationCount,
			LastFeatureTime:    p.LastFeatureTime,
			FavoritedBy:        p.FavoritedBy,
			HasCustomThumbnail: p.HasCustomThumbnail,
			MimeType:           p.MimeType,
			Comments:           p.Comments,
			Pools:              p.Pools,
		},
		Tags:      tags,
		Relations: relIds,
	}
}
