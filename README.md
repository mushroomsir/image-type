# image-type
[![Build Status](https://img.shields.io/travis/mushroomsir/image-type.svg?style=flat-square)](https://travis-ci.org/mushroomsir/image-type)
[![Coverage Status](http://img.shields.io/coveralls/mushroomsir/image-type.svg?style=flat-square)](https://coveralls.io/github/mushroomsir/image-type?branch=master)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/mushroomsir/image-type/blob/master/LICENSE)


# Installation

```sh
go get github.com/mushroomsir/image-type
```

# Usage
## parse media type
```go
package main

import (
	"fmt"
	"os"

	imageType "github.com/mushroomsir/image-type"
)

func main() {
	file, _ := os.Open("../testdata/test1.jpg")

	bytes := make([]byte, 256)
	file.Read(bytes)
	res, err := imageType.Parse(bytes)
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

# Licenses

All source code is licensed under the [MIT License](https://github.com/mushroomsir/image-type/blob/master/LICENSE).