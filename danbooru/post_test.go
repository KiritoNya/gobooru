package danbooru_test

import (
	"testing"
)

func TestPost_GetGeneralTags(t *testing.T) {
	// Get post
	p, err := Client.GetPostById(5149172)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("TAGS:", p.GetGeneralTags())
}

func TestPost_GetChatactersTags(t *testing.T) {
	// Get post
	p, err := Client.GetPostById(5149172)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("TAGS:", p.GetChatactersTags())
}

func TestPost_GetCopyrightTags(t *testing.T) {
	// Get post
	p, err := Client.GetPostById(5149172)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("TAGS:", p.GetCopyrightTags())
}

func TestPost_GetArtistTags(t *testing.T) {
	// Get post
	p, err := Client.GetPostById(5149172)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("TAGS:", p.GetArtistTags())
}

func TestPost_GetMetaTags(t *testing.T) {
	// Get post
	p, err := Client.GetPostById(5149172)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("TAGS:", p.GetMetaTags())
}