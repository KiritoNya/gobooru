package clientbooru

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"time"
)

//Client is a struct that defines the client properties
type Client struct {
	ApiKey string
	Username string
	BaseUrl string
}

//Do is a method of Client object that do a request
func (c *Client) Do(method, endpoint string, body io.Reader) ([]byte, int, error) {
	// Create client
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	// Make request
	req, err := c.MakeRequest(method, endpoint, body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// Do request
	resp, err := netClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	// Read response
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return content, resp.StatusCode, nil
}

// MakeRequest is a function that create the request with basic auth header
func (c *Client) MakeRequest(method string, endpoint string, body io.Reader) (req *http.Request, err error){
	// Check credentials
	if c.ApiKey == "" || c.Username == "" {
		return nil, errors.New("ApiKey or Username not setted")
	}

	// Create request
	req, err = http.NewRequest(
		method,
		c.joinUrl(endpoint),
		body,
	)

	// Set headers
	encodedText := base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.ApiKey))
	req.Header.Add("Authorization", "Token " + encodedText)
	req.Header.Add("Accept", "application/json")

	return
}

func (c *Client) joinUrl(endpoint string) string {
	if endpoint[0] == '/' {
		endpoint = endpoint[1:]
	}
	if c.BaseUrl[len(c.BaseUrl)-1] == '/' {
		c.BaseUrl = c.BaseUrl[:len(c.BaseUrl)]
	}

	return c.BaseUrl + "/" + endpoint
}