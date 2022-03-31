package szurubooru

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"github.com/KiritoNya/gobooru/clientbooru"
)

var Client *clientbooru.Client

// SetClient is a function that set a client for the next request
func SetClient(apiKey, username, baseUrl string) {
	Client = &clientbooru.Client{
		ApiKey:   apiKey,
		Username: username,
		BaseUrl:  baseUrl,
	}
}

// Upload is a function that upload a file and return token for the next requests
func Upload(fileData []byte, fileName string) (string, error){
	// Create multipart data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, _ := writer.CreateFormFile("content", fileName)
	io.Copy(fw, bytes.NewReader(fileData))
	writer.Close()
	// Create request
	req, err := Client.MakeRequest("POST", "/uploads", bytes.NewReader(body.Bytes()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// Do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check response code
	if resp.StatusCode != http.StatusOK  {
		return "", parseError(content)
	}

	// Read token
	raw := make(map[string]string)
	err = json.Unmarshal(content, &raw)
	if err != nil {
		return "", err
	}

	return raw["token"], nil
}

func printErr(data []byte, code int) error {
	// Parse error
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	// Create final map
	newMap := make(map[string]interface{})
	newMap["code"] = code
	newMap["error"] = raw

	// Create final json
	jsonErr, err := json.MarshalIndent(newMap, " ", "\t")
	if err != nil {
		return err
	}

	return errors.New(string(jsonErr))
}
