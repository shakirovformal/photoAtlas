package internal

import (
	"fmt"
	"os"
	"runtime"

	"github.com/disintegration/imaging"
)

func Resize(fileFF *os.File, namefile string) {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// load original image
	img, err := imaging.Open(fileFF.Name())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// resize image from 1000 to 500 while preserving the aspect ration

	// Supported resize filters: NearestNeighbor, Box, Linear, Hermite, MitchellNetravali,
	// CatmullRom, BSpline, Gaussian, Lanczos, Hann, Hamming, Blackman, Bartlett, Welch, Cosine.

	dstimg := imaging.Resize(img, 500, 0, imaging.Box)

	// save resized image
	err = imaging.Save(dstimg, "resize_"+fileFF.Name())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Ресайз прошёл и сохранён")
}
