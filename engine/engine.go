package engine

import (
	"github.com/GerryLon/go-crawler/config"
	"github.com/GerryLon/go-crawler/fetcher"
	"github.com/GerryLon/go-crawler/filter"
	"log"
	"strings"
	"time"
)

func Run(seeds ...Request) {

	var requests []Request
	requests = append(requests, seeds...)

	dedupFilter := filter.RedisDedupFilter{}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		url := r.Url

		if strings.TrimSpace(url) == "" {
			log.Println("url is empty!")
			continue
		}

		if dedupFilter.Has(url) {
			log.Printf("url %s has fetched!\n", url)
			continue
		}
		dedupFilter.Set(url)

		result := worker(r)
		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			log.Printf("Got item %v\n", item)
		}

		//// rand.Intn(n) => [0,n)
		//time.Sleep(time.Second * time.Duration(1+rand.Intn(3)))
	}
}

var rateLimiter = time.Tick(time.Second / config.QPS)

func worker(r Request) ParseResult {
	<-rateLimiter
	log.Printf("Fetching %s\n", r.Url)
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("error occured when get %s: %s\n", r.Url, err)
	}
	result := r.Parser(contents)
	return result
}
