package engine

import (
	"log"
	"strings"
)

type SimpleEngine struct {
	Deduper Deduper
}

func (e *SimpleEngine) Run(seeds ...Request) {

	var requests []Request
	//requests = append(requests, seeds...)
	for _, request := range seeds {
		if e.Deduper.isDuplicate(request.Url) {
			log.Printf("#%d: %s is duplicate", e.Deduper.Len(), request.Url)
			continue
		} else {
			requests = append(requests, request)
		}
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		url := r.Url

		if strings.TrimSpace(url) == "" {
			log.Printf("url is empty!")
			continue
		}

		result, err := Worker(r)

		if err != nil {
			continue
		}

		// requests = append(requests, result.Requests...)
		for _, request := range result.Requests {
			if e.Deduper.isDuplicate(request.Url) {
				log.Printf("#%d: %s is duplicate", e.Deduper.Len(), request.Url)
				continue
			}
			requests = append(requests, request)
		}

		for _, item := range result.Items {
			log.Printf("Got item %v\n", item)
		}

		//// rand.Intn(n) => [0,n)
		//time.Sleep(time.Second * time.Duration(1+rand.Intn(3)))
	}
}
