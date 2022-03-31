package szurubooru

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SearchResult struct {
	ExactPost *Post `json:"exact_post"`
	SimilarPosts *SimilarPost `json:"similar_posts"`
}

type SimilarPost struct {
	Distance float32 `json:"distance,omitempty"`
	Post *Post `json:"post,omitempty"`
}

// ReverseSearchByToken is a function that search if there are copy image or similar image with the token of the file uploaded.
func ReverseSearchByToken(FileToken string) (*SearchResult, error){
	// Create struct
	var raw = struct {
		ContentToken string `json:"contentToken"`
	}{
		ContentToken: FileToken,
	}

	// Create json
	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}

	// Do request
	content, code, err := Client.Do("POST", "/posts/reverse-search", bytes.NewReader(data))

	// Check status code
	if code != http.StatusOK  {
		return nil, parseError(content)
	}

	// Parse json response
	var sr SearchResult
	err = json.Unmarshal(content, &sr)
	if err != nil {
		return nil, err
	}

	return &sr, nil
}