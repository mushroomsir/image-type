package imageType

import "errors"

// ImageInfo ...
type ImageInfo struct {
	Type     string
	MimeType string
	Width    int
	Height   int
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
	if byteLen > 15 {
		img.Width = int(bytes[0])<<24 | int(bytes[1])<<16 | int(bytes[2])<<8 | int(bytes[3])
		img.Height = int(bytes[4])<<24 | int(bytes[5])<<16 | int(bytes[6])<<8 | int(bytes[7])
	}
}
func parseGif(bytes []byte, img *ImageInfo) {
	byteLen := len(bytes)
	img.MimeType = "image/gif"
	img.Type = "gif"
	if byteLen > 5 {
		img.Width = int(bytes[0]) + int(bytes[1])*256
		img.Height = int(bytes[2]) + int(bytes[3])*256
	}
}
func parseBmp(bytes []byte, img *ImageInfo) {
	byteLen := len(bytes)
	img.MimeType = "image/bmp"
	img.Type = "bmp"
	if byteLen > 17 {
		img.Width = int(bytes[3])<<24 | int(bytes[2])<<16 | int(bytes[1])<<8 | int(bytes[0])
		img.Height = int(bytes[7])<<24 | int(bytes[6])<<16 | int(bytes[5])<<8 | int(bytes[4])
	}
}
func parseWebp(bytes []byte, img *ImageInfo) {
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
