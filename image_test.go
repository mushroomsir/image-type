package imageType

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		path      string
		imageType string
		mimeType  string
		width     int
		height    int
	}{
		{"testdata/test.jpg", "jpeg", "image/jpeg", 600, 600},
		{"testdata/test.png", "png", "image/png", 612, 357},
		{"testdata/test.gif", "gif", "image/gif", 500, 500},
		{"testdata/test.webp", "webp", "image/webp", 386, 395},
		{"testdata/testLossy.webp", "webp", "image/webp", 550, 368},
		{"testdata/test.psd", "psd", "image/vnd.adobe.photoshop", 2481, 3507},
		{"testdata/test.bmp", "bmp", "image/bmp", 622, 630},
		{"testdata/test.ico", "ico", "image/x-icon", 32, 32},
		{"testdata/test-multi-size.ico", "ico", "image/x-icon", 256, 256},
	}
	for i, c := range cases {
		var res *ImageInfo
		var err error
		var file *os.File
		if i == 0 {
			res, err = ParsePath(c.path)
		} else if i == 1 {
			file, _ = os.Open(c.path)
			res, err = ParseFile(file)
		} else if i == 2 {
			file, _ = os.Open(c.path)
			res, err = ParseReader(file)
		} else {
			file, _ = os.Open(c.path)
			bytes := make([]byte, 256)
			file.Read(bytes)
			res, err = Parse(bytes)
		}
		if assert.Nil(err) {
			assert.Equal(c.imageType, res.Type)
			assert.Equal(c.mimeType, res.MimeType)
			assert.Equal(c.width, res.Width)
			assert.Equal(c.height, res.Height)
		}
		if i != 0 {
			file.Close()
		}
	}
}
func TestError(t *testing.T) {
	assert := assert.New(t)
	res, err := ParsePath("testdata/error.jpg")
	assert.NotNil(err)
	assert.Nil(res)

	res, err = Parse([]byte{'x', 'c', 'x'})
	assert.NotNil(err)
	assert.Nil(res)

	res, err = ParseReader(bytes.NewReader([]byte(strings.Repeat("x", 257))))
	assert.NotNil(err)
	assert.Nil(res)

	res, err = ParseReader(bytes.NewReader([]byte(strings.Repeat("x", 2))))
	assert.NotNil(err)
	assert.Nil(res)

	f, err := os.Create("test.txt")
	assert.Nil(err)
	res, err = ParseFile(f)
	assert.NotNil(err)
	assert.Nil(res)
	f.Close()
	os.Remove("test.txt")
}
