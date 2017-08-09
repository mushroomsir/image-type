package imageType

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJPEG(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test.jpg")
	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := Parse(bytes)
	if assert.Nil(err) {
		assert.Equal("jpg", res.Type)
		assert.Equal("image/jpeg", res.MimeType)
		assert.Equal(600, res.Width)
		assert.Equal(600, res.Height)
	}
}
func TestPng(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test.png")
	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := Parse(bytes)
	if assert.Nil(err) {
		assert.Equal("png", res.Type)
		assert.Equal("image/png", res.MimeType)
		assert.Equal(612, res.Width)
		assert.Equal(357, res.Height)
	}
}

func TestGif(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test.gif")
	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := Parse(bytes)
	if assert.Nil(err) {
		assert.Equal("gif", res.Type)
		assert.Equal("image/gif", res.MimeType)
		assert.Equal(500, res.Width)
		assert.Equal(500, res.Height)
	}
}

func TestBmp(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test.bmp")
	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := Parse(bytes)
	if assert.Nil(err) {
		assert.Equal("bmp", res.Type)
		assert.Equal("image/bmp", res.MimeType)
		assert.Equal(622, res.Width)
		assert.Equal(630, res.Height)
	}
}
