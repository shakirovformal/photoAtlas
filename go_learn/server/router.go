package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

const listenAddr = "127.0.0.1:8080"

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)

	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {

		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func handlefunc() {
	log.Println("====================")
	log.Println(ColorGreen, "SERVER IS STARTED  |", ColorReset)
	log.Println("====================")

	rout := mux.NewRouter()
	rout.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	rout.HandleFunc("/", displayHomePage).Methods("GET")
	rout.HandleFunc("/upload_image", displayUploadImage)
	rout.HandleFunc("/upload_image/1", uploadFile)
	http.Handle("/", rout)
	spa := spaHandler{staticPath: "./site/templates/", indexPath: "home_page.html"}
	rout.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler: rout,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
	//log.Fatal(srv.ListenAndServe())
	//http.ListenAndServe(listenAddr, nil)
}

func Server() {

	handlefunc()

}
