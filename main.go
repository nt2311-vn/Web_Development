package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", showInfo)
	http.HandleFunc("/site", serveFile)

	if err := http.ListenAndServe(":8999", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current time: ", time.Now())
	fmt.Fprintln(w, "URL Path: ", html.EscapeString(r.URL.Path))
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "imdex.html")
}
