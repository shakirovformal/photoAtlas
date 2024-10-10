package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func displayUploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK")
	t, err := template.ParseFiles("./site/templates/upload_image.html")
	if err != nil {
		log.Fatal("PANICCC")
	}
	t.ExecuteTemplate(w, "upl", nil)
}

func displayHomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./site/templates/home_page.html")
	if err != nil {
		log.Fatal("PANICCC")
	}
	t.Execute(w, nil)
}

func handlefunc() {
	http.HandleFunc("/home", displayHomePage)
	http.HandleFunc("/upload_image", displayUploadImage)
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":5432", nil)
}

func Server() {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	// mux := http.NewServeMux()

	// mux.HandleFunc("/", Index)

	// http.ListenAndServe(":"+port, mux)

	handlefunc()
}
