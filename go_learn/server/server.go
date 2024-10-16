package server

import (
	"fmt"
	"go_learn/cmd/app/go_learn/pkg"
	"io"
	"net/http"
	"os"
	"sync"
)

const MAX_UPLOAD_SIZE = 1073741824 // max size upload 1GB in bytes

func uploadFile(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup
	//var mu sync.Mutex

	//time_start := time.Now()
	fmt.Printf("Upload start\n")

	errs := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if errs != nil {
		fmt.Fprintln(w, errs)
		return
	}
	formdata := r.MultipartForm
	files := formdata.File["myFile"]

	for i, val := range files {
		wg.Add(2)
		if pkg.DB_Select(files[i].Filename) == true {
			return
		}
		wg.Done()
		go func() {
			pkg.DB_Insert(i, val.Filename)
			defer wg.Done()
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			out, err := os.Create("./images/" + files[i].Filename)
			defer out.Close()
			if err != nil {
				fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
				return
			}
			_, err = io.Copy(out, file) // file not files[i] !
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}()

	}
	wg.Wait()
	displayUploadedImage(w, r)
	//time_end := time.Now().Sub(time_start)
	//fmt.Printf("Was uploaded %v files\n", len(files))
	//fmt.Printf("Upload end in %v seconds\n", time_end.Seconds())

}
