package szurubooru_test

import (
	"testing"
	"github.com/KiritoNya/gobooru/szurubooru"
)

const ApiKey = "9530fcbd-c3f7-4119-8811-92eea0f7dbc5"
const Username = "KiritoNya"
const SiteUrl = "http://localhost:3000/api"

func TestTag_Create(t *testing.T) {
	// Set client
	szurubooru.SetClient(ApiKey, Username, SiteUrl)

	// Create tag
	tag := &szurubooru.Tag{
		Names:        []string{"KiritoNya"},
		Category:     "Artists",
	}

	// Create Tag
	newTag, err := tag.Create()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("TAG:", newTag)
}

func TestGetTagByName(t *testing.T) {
	// Set client
	szurubooru.SetClient(ApiKey, Username, SiteUrl)

	tag, err := szurubooru.GetTagByName("matou_sakura")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("TAG:", tag)
}
