package imageType

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJPEG(t *testing.T) {
	assert := assert.New(t)
	file, _ := os.Open("testdata/test1.jpg")

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
