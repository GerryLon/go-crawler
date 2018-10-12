package engine

import "log"
import "github.com/GerryLon/go-crawler/fetcher"

// Worker: fetch giving url and parse fetched contents
func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("error occured when get %s: %s", r.Url, err)
		return ParseResult{}, err
	}
	result := r.Parser(contents)
	return result, nil
}
