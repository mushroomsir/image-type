package imageType

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert := assert.New(t)
	res, err := ParsePath("testdata/error.jpg")
	assert.NotNil(err)
	assert.Nil(res)

	res, err = Parse([]byte{'x', 'c', 'x'})
	assert.NotNil(err)
	assert.Nil(res)

	res, err = ParseReader(bytes.NewReader([]byte{'x', 'c', 'x'}))
	assert.NotNil(err)
	assert.Nil(res)
}
func TestJPEG(t *testing.T) {
	assert := assert.New(t)
	res, err := ParsePath("testdata/test.jpg")
	if assert.Nil(err) {
		assert.Equal("jpeg", res.Type)
		assert.Equal("image/jpeg", res.MimeType)
		assert.Equal(600, res.Width)
		assert.Equal(600, res.Height)
	}
}
func TestPng(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test.png")
	res, err := ParseFile(file)
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
	res, err := ParseReader(file)
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

func TestWebp(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test.webp")
	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := Parse(bytes)
	if assert.Nil(err) {
		assert.Equal("webp", res.Type)
		assert.Equal("image/webp", res.MimeType)
		assert.Equal(386, res.Width)
		assert.Equal(395, res.Height)
	}
}

func TestWebpLossy(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/testLossy.webp")
	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := Parse(bytes)
	if assert.Nil(err) {
		assert.Equal("webp", res.Type)
		assert.Equal("image/webp", res.MimeType)
		assert.Equal(550, res.Width)
		assert.Equal(368, res.Height)
	}
}

func TestIco(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test.ico")
	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := Parse(bytes)
	if assert.Nil(err) {
		assert.Equal("ico", res.Type)
		assert.Equal("image/x-icon", res.MimeType)
	}
}

func TestPsd(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test.psd")
	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := Parse(bytes)
	if assert.Nil(err) {
		assert.Equal("psd", res.Type)
		assert.Equal("image/vnd.adobe.photoshop", res.MimeType)
		assert.Equal(2481, res.Width)
		assert.Equal(3507, res.Height)
	}
}
