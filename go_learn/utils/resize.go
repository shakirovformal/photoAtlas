package utils

import (
	"log"

	"github.com/disintegration/imaging"
)

func Resize(path, filename string) {

	src, err := imaging.Open(path + filename)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	src = imaging.Resize(src, 200, 0, imaging.Lanczos)
	err = imaging.Save(src, path+"resized/resized_"+filename)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}
