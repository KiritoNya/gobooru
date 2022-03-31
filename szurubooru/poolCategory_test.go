package szurubooru_test

import (
	"testing"
	"github.com/KiritoNya/gobooru/szurubooru"
)

func TestPoolCategory_Create(t *testing.T) {
	// Set client
	szurubooru.SetClient(ApiKey, Username, SiteUrl)

	pc := szurubooru.PoolCategory{
		Name:    "Prova",
		Color:   "#cb318e",
	}

	category, err := pc.Create()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Category: ", category)
}

func TestGetPoolCategoryByName(t *testing.T) {
	// Set client
	szurubooru.SetClient(ApiKey, Username, SiteUrl)

	category, err := szurubooru.GetPoolCategoryByName("Cosplay")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Category: ", category)
}
