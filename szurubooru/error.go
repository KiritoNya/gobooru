package szurubooru

import "encoding/json"

type Error struct {
	Name string
	Title string
	Description string
}

func (e *Error) Error() string { return e.Description }

func parseError(data []byte) error {
	var e Error
	_ = json.Unmarshal(data, &e)

	return &e
}
