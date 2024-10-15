package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func displayUploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Now displaying Upload image page")
	t, err := template.ParseFiles("./site/templates/upload_image.html")
	if err != nil {
		log.Fatal("upload image page not open")
	}
	t.Execute(w, nil)
}

func displayHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Now displaying Home page")
	t, err := template.ParseFiles("./site/templates/home_page.html")
	if err != nil {
		log.Fatal("Page home page not open")
	}
	t.Execute(w, nil)
}

func displayUploadedImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Now displaying image uploaded page")
	t, err := template.ParseFiles("./site/templates/uploaded.html")
	if err != nil {
		log.Fatal("Page uploaded files not open")
	}
	t.Execute(w, nil)

}
