package server

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

// var tpl *template.Template

// func init_uploadPage(w http.ResponseWriter, r *http.Request) {
// 	tpl = template.Must(template.ParseGlob("C:/Users/root/Desktop/photoAtlas/site/templates/upload_image.html"))
// }

// // Compile templates on start of the application
// //func init() {
// //	tpl = template.Must(template.ParseGlob("C:/Users/root/Desktop/photoAtlas/site/templates/home_page.html"))
// //}

// // Display the named template
// func display_homePage(w http.ResponseWriter, r *http.Request) {
// 	tpl.ExecuteTemplate(w, "index", nil)

// }
// func display_uploadPage(w http.ResponseWriter, r *http.Request) {
// 	tpl.ExecuteTemplate(w, "index1", nil)

// }

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
	http.ListenAndServe(":8080", nil)
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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)
	fmt.Println("OK1")

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	// Create file
	fmt.Println("OK2")
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("OK3")
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("OK4")

	//fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		displayUploadImage(w, r)
	case "POST":
		uploadFile(w, r)
	}
}
