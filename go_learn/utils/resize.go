package utils

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func Resize(path, filename string) {
	pathname := path + filename
	file, err := os.Open(pathname)
	if err != nil {
		log.Fatal(err.Error())
	}
	img, errs := png.Decode(file)
	if errs != nil {
		log.Fatal(err.Error())
	}

	file.Close()

	m := resize.Resize(200, 0, img, resize.Bicubic)

	out, err := os.Create(path + "resized_" + filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer out.Close()
	png.Encode(out, m)
	fmt.Println("File resized")

}
