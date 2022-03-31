package szurubooru

import (
	"bytes"
	"encoding/json"
	"errors"
)

type PoolCategory struct {
	Version int `json:"version,omitempty"`
	Name string `json:"name"`
	Color string `json:"color"`
	Usage int `json:"usage,omitempty"`
	Default bool `json:"default,omitempty"`
}

func GetPoolCategoryByName(name string) (*PoolCategory, error){
	// Check input
	if name == "" {
		return nil, errors.New("Name not valid")
	}

	// Do request
	respContent, statusCode, err := Client.Do("GET", "/pool-category/" + name, nil)
	if err != nil {
		return nil, err
	}

	// Check status code
	if statusCode != 200 {
		return nil, parseError(respContent)
	}

	// Parse response json
	var newPoolCategory PoolCategory
	err = json.Unmarshal(respContent, &newPoolCategory)
	if err != nil {
		return nil, err
	}

	return &newPoolCategory, nil
}

func (pc *PoolCategory) Create() (*PoolCategory, error){
	// Check input
	if pc.Name == "" || pc.Color == "" {
		return nil, errors.New("Name or color not valid")
	}

	// Create json
	data, err := json.Marshal(pc)
	if err != nil {
		return nil, err
	}

	// Do request
	respContent, statusCode, err := Client.Do("POST", "/pool-categories", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Check status code
	if statusCode != 200 {
		return nil, parseError(respContent)
	}

	// Parse response json
	var newPoolCategory PoolCategory
	err = json.Unmarshal(respContent, &newPoolCategory)
	if err != nil {
		return nil, err
	}

	return &newPoolCategory, nil
}