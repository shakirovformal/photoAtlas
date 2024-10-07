package server

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const MAX_UPLOAD_SIZE = 2048 * 2048 // 1MB

type Progress struct {
	TotalSize int64
	BytesRead int64
}

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
	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	formdata := r.MultipartForm
	files := formdata.File["myFile"]
	for i, _ := range files { // loop through the files one by one
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

		fmt.Fprintf(w, "Files uploaded successfully : ")
		fmt.Fprintf(w, files[i].Filename+"\n")

	}

	//////////////////////////////////////////////////////////

	// Create a new file in the uploads directory
	// f, err := os.OpenFile("./images/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer f.Close()

	// // Copy the contents of the file to the new file
	// _, err = io.Copy(f, file)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	//w.Write([]byte("File uploaded successfully"))
}

// // Get the file from the request
// file, handler, err := r.FormFile("myFile")
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	return
// }
// defer file.Close()

// // Create a new file in the uploads directory
// f, err := os.OpenFile("./images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }
// defer f.Close()

// // Copy the contents of the file to the new file
// _, err = io.Copy(f, file)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }

// w.Write([]byte("File uploaded successfully"))

// Get handler for filename, size and headers
// file, handler, err := r.FormFile("myFile")
// if err != nil {
// 	fmt.Println("Error Retrieving the File")
// 	fmt.Println(err)
// 	return
// }
// // Create file
// fmt.Println("OK2")
// dst, err := os.Create(handler.Filename)
// defer dst.Close()
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }
// fmt.Println("OK3")
// defer file.Close()
// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
// fmt.Printf("File Size: %+v\n", handler.Size)
// fmt.Printf("MIME Header: %+v\n", handler.Header)

// // Copy the uploaded file to the created file on the filesystem
// if _, err := io.Copy(dst, file); err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }
// fmt.Println("OK4")

// fmt.Fprintf(w, "Successfully Uploaded File\n")

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		displayUploadImage(w, r)
	case "POST":
		uploadFile(w, r)
	}
}
