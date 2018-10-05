package engine

import (
	"github.com/GerryLon/go-crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {

	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		q := requests[0]
		requests = requests[1:]

		url := q.Url
		log.Printf("Fetching %s\n", url)
		contents, err := fetcher.Fetch(url)
		if err != nil {
			log.Printf("error occured when get %s: %s\n", url, err)
		}

		result := q.Parser(contents)
		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			log.Printf("Got item %v\n", item)
		}
	}
}
