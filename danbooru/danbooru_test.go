package danbooru_test

import (
	"github.com/KiritoNya/gobooru/danbooru"
	"testing"
)

var Client = danbooru.Client{
Username: "KiritoNya",
ApiKey:   "7nkssPLAbcTkw2PGGfdVJryD",
Hostname: danbooru.SiteUrl,
}

func TestClient_GetPostById(t *testing.T) {
	// Get post
	p, err := Client.GetPostById(5149172)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("POST:", p)
}

func TestClient_GetPostByMd5(t *testing.T) {
	// Get post
	p, err := Client.GetPostByMd5("02e3e1e86653dfd9c8d6c78b72ff7a6c")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("POST:", p)
}
