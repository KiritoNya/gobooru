package szurubooru

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type Tag struct {
	Version      int      `json:"version,omitempty"`
	Names        []string      `json:"names"`
	Category     string      `json:"category"`
	Implications []*MicroTag `json:"implications,omitempty"`
	Suggestions  []*MicroTag `json:"suggestions,omitempty"`
	CreationTime *time.Time  `json:"creation_time,omitempty"`
	LastEditTime *time.Time  `json:"last_edit_time,omitempty"`
	UsageCount   int         `json:"usage_count,omitempty"`
	Description  string      `json:"description,omitempty"`
}

type MicroTag struct {
	Names    []string `json:"names"`
	Category string `json:"category"`
	Usages   int	`json:"usages,omitempty"`
}

func GetTagByName(name string) (*Tag, error) {
	// Do request
	resp, code, err := Client.Do("GET", "/tag/" + name, nil)

	// La risorsa non esiste
	if code != 200 {
		if code == 404 {
			return nil, nil
		}
		return nil, parseError(resp)
	}

	// Check error
	if err != nil {
		return nil, err
	}

	// Parse json response
	var t Tag
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Tag) Create() (*Tag, error) {
	// Check validity
	if t.Names == nil || t.Category == "" {
		return nil, errors.New("Names and Category are required")
	}

	// Create Json
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	// Do request
	respContent, statusCode, err := Client.Do("POST", "/tags", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Check status code
	if statusCode != 200 {
		return nil, parseError(respContent)
	}

	// Parse response json
	var newTag Tag
	err = json.Unmarshal(respContent, &newTag)
	if err != nil {
		return nil, err
	}

	return &newTag, nil
}