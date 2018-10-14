package main

import (
	"github.com/GerryLon/go-crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/search",
		controller.CreateSearchResultHandler("go-crawler/frontend/view/index.html"))
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
