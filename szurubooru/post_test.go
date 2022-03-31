package szurubooru_test

import (
	"os"
	"testing"
	"github.com/KiritoNya/gobooru/szurubooru"
)

func TestPost_Create(t *testing.T) {
	// Set client
	szurubooru.SetClient(ApiKey, Username, SiteUrl)

	post := &szurubooru.Post{
		Safety: "safe",
		Tags: []*szurubooru.MicroTag{
			{
				Names:    []string{"KiritoNya"},
				Category: "artist",
			},
			{
				Names:    []string{"Sergio"},
				Category: "character",
			},
		},
	}

	content, err := os.ReadFile("80307584-3a54c900-87ca-11ea-9aee-920069e59979.png")
	if err != nil {
		t.Fatal(err)
	}

	post, err = post.Create(content, "photo_2022-02-21_19-58-10.png")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("POST:", post)
}
func TestPost_Update(t *testing.T) {
	// Set client
	szurubooru.SetClient(ApiKey, Username, SiteUrl)

	post := &szurubooru.Post{
		Id: 957,
		Version: 1,
		Safety:             "safe",
		Relations: []*szurubooru.Post{
			{Id: 948},
		},
	}

	newPost, err := post.Update()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(newPost)
}

func TestPost_Delete(t *testing.T) {
	// Set client
	szurubooru.SetClient(ApiKey, Username, SiteUrl)

	p := &szurubooru.Post{Id: 956, Version: 2}

	err := p.Delete()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Post Delete [OK]")
}
