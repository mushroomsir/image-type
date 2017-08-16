package imageType

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// ImageInfo ...
type ImageInfo struct {
	Type     string
	MimeType string
	Width    int
	Height   int
}

var (
	headerLength     = 256
	ErrInvalidImage  = errors.New("invalid image")
	ErrInvalidLength = errors.New("invalid length")
)

// ParsePath ...
func ParsePath(filePath string) (img *ImageInfo, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes := make([]byte, headerLength)
	file.Read(bytes)
	return Parse(bytes)
}

// ParseFile ...
func ParseFile(file *os.File) (img *ImageInfo, err error) {
	bytes := make([]byte, headerLength)
	_, err = file.Read(bytes)
	if err != nil {
		return
	}
	return Parse(bytes)
}

// ParseReader ...
func ParseReader(rd io.Reader) (img *ImageInfo, err error) {
	br := bufio.NewReader(rd)
	bytes, err := br.Peek(headerLength)
	if err != nil {
		return
	}
	return Parse(bytes)
}

// Parse ...
func Parse(bytes []byte) (img *ImageInfo, err error) {
	img = &ImageInfo{}
	byteLen := len(bytes)
	if byteLen < headerLength {
		return nil, ErrInvalidLength
	}
	if bytes[0] == 0xFF && bytes[1] == 0xD8 && bytes[2] == 0xFF {
		parseJpg(bytes, img)
	} else if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
		parsePng(bytes, img)
	} else if bytes[0] == 0x47 && bytes[1] == 0x49 && bytes[2] == 0x46 {
		parseGif(bytes, img)
	} else if bytes[0] == 0x42 && bytes[1] == 0x4D {
		parseBmp(bytes, img)
	} else if bytes[8] == 0x57 && bytes[9] == 0x45 && bytes[10] == 0x42 && bytes[11] == 0x50 {
		parseWebp(bytes, img)
	} else if bytes[0] == 0x00 && bytes[1] == 0x00 && bytes[2] == 0x01 && bytes[3] == 0x00 {
		parseIco(bytes, img)
	} else if bytes[0] == 0x38 && bytes[1] == 0x42 && bytes[2] == 0x50 && bytes[3] == 0x53 {
		parsePsd(bytes, img)
	} else if binary.BigEndian.Uint32(bytes) == 0x49492a00 || binary.BigEndian.Uint32(bytes) == 0x4d4d002a {
		parseTiff(bytes, img)
	}
	if len(img.Type) == 0 {
		img = nil
		err = ErrInvalidImage
	}
	return
}

func parseJpg(bytes []byte, img *ImageInfo) {
	byteLen := len(bytes)
	img.MimeType = "image/jpeg"
	img.Type = "jpeg"
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
			break
		}
		position += 2
	}
	return
}
func parsePng(bytes []byte, img *ImageInfo) {
	img.MimeType = "image/png"
	img.Type = "png"
	r := bytes[16:]
	img.Width = int(r[0])<<24 | int(r[1])<<16 | int(r[2])<<8 | int(r[3])
	img.Height = int(r[4])<<24 | int(r[5])<<16 | int(r[6])<<8 | int(r[7])
}
func parseGif(bytes []byte, img *ImageInfo) {
	img.MimeType = "image/gif"
	img.Type = "gif"
	r := bytes[6:]
	img.Width = int(r[0]) + int(r[1])*256
	img.Height = int(r[2]) + int(r[3])*256
}

func parseBmp(bytes []byte, img *ImageInfo) {
	img.MimeType = "image/bmp"
	img.Type = "bmp"
	r := bytes[18:]
	img.Width = int(r[3])<<24 | int(r[2])<<16 | int(r[1])<<8 | int(r[0])
	img.Height = int(r[7])<<24 | int(r[6])<<16 | int(r[5])<<8 | int(r[4])
}

/**
webp
https://chromium.googlesource.com/webm/libwebp/+/master/doc/webp-lossless-bitstream-spec.txt
https://developers.google.com/speed/webp/docs/riff_container

*/
func parseWebp(bytes []byte, img *ImageInfo) {
	img.MimeType = "image/webp"
	img.Type = "webp"
	r := bytes[12:]
	if r[0] == 0x56 && r[1] == 0x50 && r[2] == 0x38 && r[3] == 0x4C {
		r = bytes[21:]
		img.Width = 1 + (((int(r[1]) & 0x3F) << 8) | int(r[1]))
		img.Height = 1 + (((int(r[3]) & 0xF) << 10) | int(r[2])<<2 | (int(r[1]) & 0xC0 >> 6))
	} else {
		// https://tools.ietf.org/html/rfc6386#section-9
		r = bytes[23:]
		if r[0] == 0x9d && r[1] == 0x01 && r[2] == 0x2a {
			img.Width = int(r[4]&0x3f)<<8 | int(r[3])
			img.Height = int(r[6]&0x3f)<<8 | int(r[5])
		}
	}
}

/*
 * ICON Header
 *
 * | Offset | Size | Purpose                                                                                   |
 * | 0	    | 2    | Reserved. Must always be 0.                                                               |
 * | 2      | 2    | Image type: 1 for icon (.ICO) image, 2 for cursor (.CUR) image. Other values are invalid. |
 * | 4      | 2    | Number of images in the file.                                                             |
 *
 */
var sizeHeader = 2 + 2 + 2 // 6
/*
 * Image Entry
 *
 * | Offset | Size | Purpose                                                                                          |
 * | 0	    | 1    | Image width in pixels. Can be any number between 0 and 255. Value 0 means width is 256 pixels.   |
 * | 1      | 1    | Image height in pixels. Can be any number between 0 and 255. Value 0 means height is 256 pixels. |
 * | 2      | 1    | Number of colors in the color palette. Should be 0 if the image does not use a color palette.    |
 * | 3      | 1    | Reserved. Should be 0.                                                                           |
 * | 4      | 2    | ICO format: Color planes. Should be 0 or 1.                                                      |
 * |        |      | CUR format: The horizontal coordinates of the hotspot in number of pixels from the left.         |
 * | 6      | 2    | ICO format: Bits per pixel.                                                                      |
 * |        |      | CUR format: The vertical coordinates of the hotspot in number of pixels from the top.            |
 * | 8      | 4    | The size of the image's data in bytes                                                            |
 * | 12     | 4    | The offset of BMP or PNG data from the beginning of the ICO/CUR file                             |
 *
 */
var sizeImageEntry = 1 + 1 + 1 + 1 + 2 + 2 + 4 + 4 // 16

// Just extract dimension from the first image
func parseIco(bytes []byte, img *ImageInfo) {
	img.MimeType = "image/x-icon"
	img.Type = "ico"
	img.Width = int(bytes[sizeHeader])
	img.Height = int(bytes[sizeHeader+1])
	if img.Width == 0 {
		img.Width = 256
	}
	if img.Height == 0 {
		img.Height = 256
	}
}

func parsePsd(bytes []byte, img *ImageInfo) {
	img.MimeType = "image/vnd.adobe.photoshop"
	img.Type = "psd"
	img.Width = int(binary.BigEndian.Uint32(bytes[18:]))
	img.Height = int(binary.BigEndian.Uint32(bytes[14:]))
}

func parseTiff(bytes []byte, img *ImageInfo) {
	img.MimeType = "image/tiff"
	img.Type = "tiff"
}
