package danbooru

import "strings"

type Post struct {
	ID                  int         `json:"id"`
	CreatedAt           string      `json:"created_at"`
	UploaderID          int         `json:"uploader_id"`
	Score               int         `json:"score"`
	Source              string      `json:"source"`
	Md5                 string      `json:"md5"`
	LastCommentBumpedAt interface{} `json:"last_comment_bumped_at"`
	Rating              string      `json:"rating"`
	ImageWidth          int         `json:"image_width"`
	ImageHeight         int         `json:"image_height"`
	TagString           string      `json:"tag_string"`
	FavCount            int         `json:"fav_count"`
	FileExt             string      `json:"file_ext"`
	LastNotedAt         string      `json:"last_noted_at"`
	ParentID            interface{} `json:"parent_id"`
	HasChildren         bool        `json:"has_children"`
	ApproverID          int         `json:"approver_id"`
	TagCountGeneral     int         `json:"tag_count_general"`
	TagCountArtist      int         `json:"tag_count_artist"`
	TagCountCharacter   int         `json:"tag_count_character"`
	TagCountCopyright   int         `json:"tag_count_copyright"`
	FileSize            int         `json:"file_size"`
	UpScore             int         `json:"up_score"`
	DownScore           int         `json:"down_score"`
	IsPending           bool        `json:"is_pending"`
	IsFlagged           bool        `json:"is_flagged"`
	IsDeleted           bool        `json:"is_deleted"`
	TagCount            int         `json:"tag_count"`
	UpdatedAt           string      `json:"updated_at"`
	IsBanned            bool        `json:"is_banned"`
	PixivID             int         `json:"pixiv_id"`
	LastCommentedAt     interface{} `json:"last_commented_at"`
	HasActiveChildren   bool        `json:"has_active_children"`
	BitFlags            int         `json:"bit_flags"`
	TagCountMeta        int         `json:"tag_count_meta"`
	HasLarge            bool        `json:"has_large"`
	HasVisibleChildren  bool        `json:"has_visible_children"`
	TagStringGeneral    string      `json:"tag_string_general"`
	TagStringCharacter  string      `json:"tag_string_character"`
	TagStringCopyright  string      `json:"tag_string_copyright"`
	TagStringArtist     string      `json:"tag_string_artist"`
	TagStringMeta       string      `json:"tag_string_meta"`
	FileURL             string      `json:"file_url"`
	LargeFileURL        string      `json:"large_file_url"`
	PreviewFileURL      string      `json:"preview_file_url"`
}

// GetGeneralTags is a method of Post object that return a lis of general tags
func (p *Post) GetGeneralTags() []string {
	return strings.Split(p.TagStringGeneral, " ")
}

// GetChatactersTags is a method of Post object that return a lis of characters tags
func (p *Post) GetChatactersTags() []string {
	return strings.Split(p.TagStringCharacter, " ")
}

// GetCopyrightTags is a method of Post object that return a lis of copyrights tags
func (p *Post) GetCopyrightTags() []string {
	return strings.Split(p.TagStringCopyright, " ")
}

// GetArtistTags is a method of Post object that return a lis of artists tags
func (p *Post) GetArtistTags() []string {
	return strings.Split(p.TagStringArtist, " ")
}

// GetMetaTags is a method of Post object that return a lis of meta tags
func (p *Post) GetMetaTags() []string {
	return strings.Split(p.TagStringMeta, " ")
}