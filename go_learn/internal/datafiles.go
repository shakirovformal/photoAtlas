package internal

import (
	"os"
)

type Database struct {
	host string
}

// Функция которая проверяет наличие файла в директории
func Get_info_inDir(filename string) bool {
	path := "././images/"

	return searchFileinfo(path, filename)

}

func searchFileinfo(path, filename string) bool {
	lst, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var res bool
	for _, val := range lst {
		if filename == val.Name() {
			res = true
			return res
		} else if filename != val.Name() {
			res = false
		}
	}

	return res
}
