package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// http.HandleFunc("/", showInfo)
	//
	// files := http.FileServer(http.Dir(filepath.Join("var", "www")))
	// http.Handle("/site/", http.StripPrefix("/site/", fils))

	// finalHandler := http.HandlerFunc(final)
	// http.Handle("/", middleware1(middleware2(finalHandler)))

	router := mux.NewRouter()
	router.HandleFunc("/product/{id:[0-9]+}", pageHandler)

	http.Handle("/", router)

	if err := http.ListenAndServe(":8999", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]
	log.Printf("Product ID: %s\n", productID)

	fileName := productID + ".html"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Println("No such product")
		fileName = "invalid.html"
	}

	http.ServeFile(w, r, filepath.Join("var", "www", "product", fileName))
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current time: ", time.Now())
	fmt.Fprintln(w, "URL Path: ", html.EscapeString(r.URL.Path))
}
