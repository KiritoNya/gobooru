package szurubooru

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

// Pool is a struct with the pool's properties
type Pool struct {
	Version int `json:"version,omitempty"`
	Id int `json:"id,omitempty"`
	Names []string `json:"names"`
	Category string `json:"category"`
	Posts []*Post `json:"posts,omitempty"`
	CreationTime time.Time `json:"creationTime"`
	LastEditTime time.Time `json:"lastEditTime"`
	PostCount int `json:"postCount"`
	Description string `json:"description"`
}

func (p *Pool) Create() (*Pool, error){
	// Check name
	if p.Names == nil {
		return nil, errors.New("Pool's name not setted")
	}

	// Check category
	if p.Category == "" {
		return nil, errors.New("Pool's category name not setted")
	}

	if _, err := GetPoolCategoryByName(p.Category); err != nil {
		return nil, errors.New("Pool's category not found")
	}

	// Create json
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// Do request
	respContent, statusCode, err := Client.Do("POST", "/pool", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Check status code
	if statusCode != 200 {
		return nil, parseError(respContent)
	}

	// Parse response json
	var newPool Pool
	err = json.Unmarshal(respContent, &newPool)
	if err != nil {
		return nil, err
	}

	return &newPool, nil
}
