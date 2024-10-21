package server

import (
	"fmt"
	"go_learn/cmd/app/go_learn/internal"
	"go_learn/cmd/app/go_learn/pkg"
	"go_learn/cmd/app/go_learn/utils"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

const ColorReset = string("\033[0m")
const ColorRed = string("\033[31m")
const ColorGreen = string("\033[32m")
const ColorYellow = string("\033[33m")
const ColorBlue = string("\033[34m")
const ColorPurple = string("\033[35m")
const ColorCyan = string("\033[36m")
const ColorWhite = string("\033[37m")
const MAX_UPLOAD_SIZE = 1073741824 // max size upload 1GB in bytes

func uploadFile(w http.ResponseWriter, r *http.Request) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	path := "./images/"
	time_start := time.Now()
	fmt.Printf("Upload start\n")

	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	formdata := r.MultipartForm
	files := formdata.File["selectedFile"]
	counter := []int{}
	for i, _ := range files {

		// Запускаем проверку на наличие файла в директории и БД

		wg.Add(1)
		go func() {
			defer wg.Done()
			// Проверяем наличие файла в директории
			switch internal.Get_info_inDir(files[i].Filename) {
			case true:
				switch pkg.DB_Select(files[i].Filename) {
				case 1:
					// 0.jpg ok
					fmt.Println("File ", files[i].Filename, " найден в директории и в базе, скачивать не будем")
				case 0:
					// 1.jpg ok
					mu.Lock()
					pkg.DB_Insert(files[i].Filename)
					mu.Unlock()
					fmt.Println("File ", files[i].Filename, " найден в директории но не в базе, поэтому добавим в базу значение о файле")
				}
			case false:
				switch pkg.DB_Select(files[i].Filename) {
				case 1:
					// 3.jpg ok
					mu.Lock()
					pkg.DB_UpdateNameFile(files[i].Filename)
					uploads(i, w, r)
					pkg.DB_Insert(files[i].Filename)
					mu.Unlock()
					fmt.Println("File ", files[i].Filename, " не найден в директории, но найден в базе, поэтому сделаем update с пометкой unreachable и скачаем новый файл")

				case 0:
					// 2.jpg скачать
					mu.Lock()
					pkg.DB_Insert(files[i].Filename)
					uploads(i, w, r)
					mu.Unlock()

					fmt.Println("File ", files[i].Filename, " не найден в директории и в базе, поэтому скачаем файл")
				}

			}
			utils.Resize(path, files[i].Filename)
		}()

	}
	wg.Wait()

	time_end := time.Now().Sub(time_start)
	fmt.Printf("Upload end in %v seconds\n", time_end.Seconds())
	displayUploadedImage(counter, w, r)
}

func uploads(iter int, w http.ResponseWriter, r *http.Request) {
	files := r.MultipartForm.File["selectedFile"]
	for iter, _ = range files {
		file, err := files[iter].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		out, err := os.Create("./images/" + files[iter].Filename)
		defer out.Close()
		if err != nil {
			fmt.Fprintf(w, "Unable for create, check privileges")
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
	}

}
