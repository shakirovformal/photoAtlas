package internal

import (
	"fmt"
	"io/ioutil"
	"log"
)

type DataFile struct {
	File_name     string
	File_size     int
	File_location string
}

func (idf DataFile) FName() string {
	return idf.File_name
}

func (idf DataFile) FSize() int {
	return idf.File_size
}

func (idf DataFile) FLocation() string {
	return idf.File_location
}
func (idf DataFile) getAllInfo() struct{} {
	return struct{}{}
}

// Функция которая сама узнаёт всё о файле посредством информации о имени файла
// TODO Сделать реализацию поиска файла и создания записи полученной информации в структуру DataFile
func Get_info(filename string) {

	path := "././images/"

	trace := searchFileinfo(path, filename)

	fmt.Println(trace)

}

func searchFileinfo(path, filename string) bool {
	fmt.Println("look is", filename)
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.Name() == filename {
			fmt.Println("file is found")
			return true
		} else {
			return false
			log.Fatal(err)
		}
	}
	return false
}
