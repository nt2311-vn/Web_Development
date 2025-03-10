package main

import (
	"net/http"
	"path/filepath"
)

func main() {
	http.ListenAndServe(":8999", http.FileServer(http.Dir(filepath.Join("var", "www"))))
}
