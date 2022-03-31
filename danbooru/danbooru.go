package danbooru

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

const SiteUrl string = "https://danbooru.donmai.us/"

type Client struct {
	Username string
	ApiKey string
	Hostname string
	http http.Client
}

//GetPostById is a function that get a post by id
func (c *Client) GetPostById(id int) (*Post, error) {
	var p Post

	// Do request
	resp, err := c.doRequest("GET", c.Hostname + "/posts/" + strconv.Itoa(id) + ".json", nil)
	if err != nil {
		return nil, err
	}

	// Parse json
	err = json.Unmarshal(resp, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

//GetPostByMd5 is a function that get a post by id
func (c *Client) GetPostByMd5(md5 string) (*Post, error) {
	var p Post

	// Do request
	resp, err := c.doRequest("GET", c.Hostname + "/posts/" + md5 + ".json", nil)
	if err != nil {
		return nil, err
	}

	// Parse json
	err = json.Unmarshal(resp, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// makeRequest make a request from client attributes
func (c *Client) makeRequest(method, endpoint string, body io.Reader) (req *http.Request, err error) {
	// Create Request
	req, err = http.NewRequest(
		method,
		endpoint,
		body,
	)
	if err != nil {
		return
	}

	// Add basic auth
	req.SetBasicAuth(c.Username, c.ApiKey)

	return
}

//doRequest do a request
func (c *Client) doRequest(method, endpoint string, body io.Reader) ([]byte, error) {
	// Make request
	req, err := c.makeRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	// Do request
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
