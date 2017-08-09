package main

import (
	"fmt"

	imageType "github.com/mushroomsir/image-type"
)

func main() {
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
