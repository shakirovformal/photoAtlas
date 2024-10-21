package utils

import (
	"fmt"
	"image/jpeg"
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
	img, errs := jpeg.Decode(file)
	if errs != nil {
		log.Fatal(err.Error())
		recover()
	}

	file.Close()

	m := resize.Resize(200, 0, img, resize.Bicubic)

	out, err := os.Create(path + "resized/resized_" + filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
	fmt.Println("File resized")
}
