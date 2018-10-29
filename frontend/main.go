package main

import (
	"github.com/GerryLon/go-crawler/frontend/controller"
	"net/http"
)

func main() {
	// serve static
	http.Handle("/", http.FileServer(http.Dir("go-crawler/frontend")))

	http.Handle("/search",
		controller.CreateSearchResultHandler("go-crawler/frontend/view/index.html"))
	// err := http.ListenAndServeTLS(":8888", "go-crawler/cert.pem", "go-crawler/key.pem", nil)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
