package imageType

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	cases := []struct {
		path      string
		imageType string
		mimeType  string
		width     int
		height    int
	}{
		{"testdata/test2.jpg", "jpeg", "image/jpeg", 198, 169},
		{"testdata/test.jpg", "jpeg", "image/jpeg", 600, 600},
		{"testdata/test.png", "png", "image/png", 612, 357},
		{"testdata/test.gif", "gif", "image/gif", 500, 500},
		{"testdata/test.webp", "webp", "image/webp", 386, 395},
		{"testdata/testLossy.webp", "webp", "image/webp", 550, 368},
		{"testdata/test.psd", "psd", "image/vnd.adobe.photoshop", 2481, 3507},
		{"testdata/test.bmp", "bmp", "image/bmp", 622, 630},
		{"testdata/test.ico", "ico", "image/x-icon", 32, 32},
		{"testdata/test-multi-size.ico", "ico", "image/x-icon", 256, 256},
		{"testdata/test.tiff", "tiff", "image/tiff", 1600, 2100},
		{"testdata/test.dds", "dds", "image/vnd-ms.dds", 123, 456},
	}
	for i, c := range cases {
		var res *ImageInfo
		var err error
		var file *os.File
		if i == 0 {
			res, err = ParsePath(c.path)
		} else {
			file, _ = os.Open(c.path)
			res, err = Parse(file)
		}
		require.Nil(err)
		assert.Equal(c.imageType, res.Type)
		assert.Equal(c.mimeType, res.MimeType)
		assert.Equal(c.width, res.Width)
		assert.Equal(c.height, res.Height)
		if i != 0 {
			file.Close()
		}
	}
}
func TestError(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		n int
	}{
		{2},
		{3},
		{7},
		{11},
		{257},
	}
	for _, n := range cases {
		r := bytes.NewReader([]byte(strings.Repeat("x", n.n)))
		res, err := Parse(r)
		assert.NotNil(err)
		assert.Nil(res)
	}

	res, err := ParsePath("testdata/error.jpg")
	assert.NotNil(err)
	assert.Nil(res)

	f, err := os.Create("test.txt")
	assert.Nil(err)
	res, err = Parse(f)
	assert.NotNil(err)
	assert.Nil(res)
	f.Close()
	os.Remove("test.txt")
}

func TestError2(t *testing.T) {
	assert := assert.New(t)
	case2 := []struct {
		n []byte
	}{
		{[]byte{0x89, 0x50, 0x4E, 0x47, 0x00}},
		{[]byte{0x47, 0x49, 0x46, 0x00}},
		{[]byte{0x42, 0x4D, 0x00}},
		{[]byte{0x00, 0x00, 0x01, 0x00}},
		{[]byte{0x38, 0x42, 0x50, 0x53}},
		{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x50, 0x4E, 0x47, 0x57, 0x45, 0x42, 0x50}},
		{[]byte{0x44, 0x44, 0x53, 0x20}},
	}
	for _, n := range case2 {
		r := bytes.NewReader(n.n)
		res, err := Parse(r)
		assert.NotNil(err)
		assert.Nil(res)
	}
}
