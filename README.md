# image-type
[![Build Status](https://img.shields.io/travis/mushroomsir/image-type.svg?style=flat-square)](https://travis-ci.org/mushroomsir/image-type)
[![Coverage Status](http://img.shields.io/coveralls/mushroomsir/image-type.svg?style=flat-square)](https://coveralls.io/github/mushroomsir/image-type?branch=master)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/mushroomsir/image-type/blob/master/LICENSE)


# Installation

```sh
go get github.com/mushroomsir/image-type
```

# Support 
| Format | Type    | MimeType | Dimension |
| ----- | ------- | -------- | --------- |
| jpg   | support | support  | support   |
| png   | support | support  | support   |
| gif   | support | support  | support   |
| bmp   | support | support  | support   |
| webp  | support | support  | support   |
| webp(lossy)  | support | support  | support   |
| ico   | support | support  | no        |



# Usage
## parse image
```go
package main

import (
	"fmt"

	imageType "github.com/mushroomsir/image-type"
)

func main() {
	// imageType.ParseFile(file *os.File)
	// imageType.ParseReader(rd io.Reader)
	// imageType.Parse(bytes []byte)
	res, err := imageType.ParsePath("../testdata/test.jpg")
	if err == nil {
		fmt.Println(res.Type)     // jpg
		fmt.Println(res.MimeType) // image/jpeg
		fmt.Println(res.Width)    // 600
		fmt.Println(res.Height)   // 600
	} else {
		fmt.Println(err)
	}
}
```
## check image

```go
res, _ := imageType.ParsePath("../testdata/test.jpg")
if res != nil {
	fmt.Println("It's image")
}
```

# Licenses

All source code is licensed under the [MIT License](https://github.com/mushroomsir/image-type/blob/master/LICENSE).
