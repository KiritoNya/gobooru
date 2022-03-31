package szurubooru_test

import (
	"testing"
	"time"
	"github.com/KiritoNya/gobooru/szurubooru"
)

func TestPool_Create(t *testing.T) {
	// Set client
	szurubooru.SetClient(ApiKey, Username, SiteUrl)

	p := &szurubooru.Pool{
		Names:        []string {"ProvaPool"},
		Category:     "Cosplay",
		Posts:        nil,
		CreationTime: time.Now(),
	}

	newPool, err := p.Create()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Category: ", newPool)
}
