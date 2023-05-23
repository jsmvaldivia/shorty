package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetShortenedUrl(t *testing.T) {
	bigUrl := "https://go.dev/doc/tutorial/add-a-test"

	service := ShortenerService{
		Shortener:  NewRandomShortener(),
		Repository: NewInMemoryRepository(),
	}
	shortUrl, _ := service.CreateShortUrl(bigUrl)

	assert.Equal(t, "", shortUrl, "They should be equal")
}

func TestGetShortenedUrlWithEmptyUrl(t *testing.T) {
	bigUrl := ""

	service := ShortenerService{Shortener: NewRandomShortener()}
	shortUrl, _ := service.CreateShortUrl(bigUrl)

	assert.Equal(t, "", shortUrl, "They should be equal")
}
