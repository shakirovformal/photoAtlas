package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

const MAX_UPLOAD_SIZE = 1073741824 // max size upload 1GB in bytes

func uploadFile(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	time_start := time.Now()
	fmt.Printf("Upload start\n")

	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	formdata := r.MultipartForm
	files := formdata.File["myFile"]

	for i, _ := range files {
		wg.Add(1)
		go func() {
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
			_, err = io.Copy(out, file)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}()
	}
	wg.Wait()

	time_end := time.Now().Sub(time_start)
	fmt.Printf("Was uploaded %v files\n", len(files))
	fmt.Printf("Upload end in %v seconds\n", time_end.Seconds())

}
