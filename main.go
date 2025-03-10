package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func main() {
	http.HandleFunc("/", showInfo)

	files := http.FileServer(http.Dir(filepath.Join("var", "www")))
	http.Handle("/site/", http.StripPrefix("/site/", files))

	if err := http.ListenAndServe(":8999", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current time: ", time.Now())
	fmt.Fprintln(w, "URL Path: ", html.EscapeString(r.URL.Path))
}
