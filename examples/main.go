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
