package engine

import (
	"github.com/GerryLon/go-crawler/fetcher"
	"github.com/GerryLon/go-crawler/filter"
	"log"
	"math/rand"
	"strings"
	"time"
)

func Run(seeds ...Request) {

	var requests []Request
	requests = append(requests, seeds...)

	dedupFilter := filter.RedisDedupFilter{}

	for len(requests) > 0 {
		q := requests[0]
		requests = requests[1:]

		url := q.Url

		if strings.TrimSpace(url) == "" {
			log.Println("url is empty!")
			continue
		}

		if dedupFilter.Has(url) {
			log.Printf("url %s has fetched!\n", url)
			continue
		}
		dedupFilter.Set(url)

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

		// rand.Intn(n) => [0,n)
		time.Sleep(time.Second * time.Duration(1+rand.Intn(3)))
	}
}
