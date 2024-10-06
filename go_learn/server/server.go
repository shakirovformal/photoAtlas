package server

import (
	"html/template"
	"net/http"
)

func upload_page(w http.ResponseWriter, r *http.Request) {
	tmm, _ := template.ParseFiles("./site/templates/upload_image.html")
	tmm.Execute(w, nil)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	tmm, _ := template.ParseFiles("./site/templates/home_page.html")
	tmm.Execute(w, nil)
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/upload_image", upload_page)
	http.ListenAndServe(":8080", nil)
}

func Server() {
	handleRequest()
}
