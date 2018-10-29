package main

import (
	"github.com/GerryLon/go-crawler/frontend/controller"
	"github.com/GerryLon/go-crawler/utils"
	"net/http"
	"path/filepath"
)

// http://localhost:8888/search
func main() {

	projAbsPath := utils.GetProjAbsPath(
		"github.com", "GerryLon", "go-crawler")

	// serve static
	http.Handle("/", http.FileServer(http.Dir(
		filepath.Join(projAbsPath, "frontend"))))

	http.Handle("/search",
		controller.CreateSearchResultHandler(
			filepath.Join(projAbsPath, "frontend", "view", "index.html")))
	// err := http.ListenAndServeTLS(":8888", "go-crawler/cert.pem", "go-crawler/key.pem", nil)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
