package imageType

import (
	"bufio"
	"errors"
	"io"
)

// ImageInfo ...
type ImageInfo struct {
	Type     string
	MimeType string
	Width    int
	Height   int
}

// ParseReader ...
func ParseReader(rd io.Reader) (img *ImageInfo, err error) {
	br := bufio.NewReader(rd)
	bytes, err := br.Peek(256)
	if err != nil {
		return
	}
	return Parse(bytes)
}

// Parse ...
func Parse(bytes []byte) (img *ImageInfo, err error) {
	img = &ImageInfo{}
	byteLen := len(bytes)
	if byteLen > 2 && bytes[0] == 0xFF && bytes[1] == 0xD8 && bytes[2] == 0xFF {
		parseJpg(bytes, img)
	} else if byteLen > 3 && bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
		parsePng(bytes, img)
	} else if byteLen > 2 && bytes[0] == 0x47 && bytes[1] == 0x49 && bytes[2] == 0x46 {
		parseGif(bytes, img)
	} else if byteLen > 2 && bytes[0] == 0x42 && bytes[1] == 0x4D {
		parseBmp(bytes, img)
	} else if byteLen > 11 && bytes[8] == 0x57 && bytes[9] == 0x45 && bytes[10] == 0x42 && bytes[11] == 0x50 {
		parseWebp(bytes, img)
	} else if byteLen > 3 && bytes[0] == 0x00 && bytes[1] == 0x00 && bytes[2] == 0x01 && bytes[3] == 0x00 {
		parseIco(bytes, img)
	}
	if len(img.Type) == 0 {
		img = nil
		err = errors.New("invalid image")
	}
	return
}

func parseJpg(bytes []byte, img *ImageInfo) {
	byteLen := len(bytes)
	img.MimeType = "image/jpeg"
	img.Type = "jpg"
	if byteLen < 6 {
		return
	}
	position := int64(4)
	r := bytes[position:]
	length := int(r[0]<<8) + int(r[1])
	for position < int64(byteLen) {
		position += int64(length)
		r = bytes[position:]
		length = int(r[2])<<8 + int(r[3])
		if (r[1] == 0xC0 || r[1] == 0xC2) && r[0] == 0xFF && length > 7 {
			r = bytes[position+5:]
			img.Width = int(r[2])<<8 + int(r[3])
			img.Height = int(r[0])<<8 + int(r[1])
			return
		}
		position += 2
	}
}
func parsePng(bytes []byte, img *ImageInfo) {
	byteLen := len(bytes)
	img.MimeType = "image/png"
	img.Type = "png"
	if byteLen > 23 {
		r := bytes[16:]
		img.Width = int(r[0])<<24 | int(r[1])<<16 | int(r[2])<<8 | int(r[3])
		img.Height = int(r[4])<<24 | int(r[5])<<16 | int(r[6])<<8 | int(r[7])
	}
}
func parseGif(bytes []byte, img *ImageInfo) {
	byteLen := len(bytes)
	img.MimeType = "image/gif"
	img.Type = "gif"
	if byteLen > 5 {
		r := bytes[6:]
		img.Width = int(r[0]) + int(r[1])*256
		img.Height = int(r[2]) + int(r[3])*256
	}
}

func parseBmp(bytes []byte, img *ImageInfo) {
	byteLen := len(bytes)
	img.MimeType = "image/bmp"
	img.Type = "bmp"
	if byteLen > 21 {
		r := bytes[18:]
		img.Width = int(r[3])<<24 | int(r[2])<<16 | int(r[1])<<8 | int(r[0])
		img.Height = int(r[7])<<24 | int(r[6])<<16 | int(r[5])<<8 | int(r[4])
	}
}
func parseWebp(data []byte, img *ImageInfo) {
	img.MimeType = "image/webp"
	img.Type = "webp"
}

func parseIco(bytes []byte, img *ImageInfo) {
	img.MimeType = "image/x-icon"
	img.Type = "ico"
}

// IsImage ...
func IsImage(bytes []byte) bool {
	img, _ := Parse(bytes)
	return img != nil
}
