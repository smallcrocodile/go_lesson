package engine

import (
	"crawer/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	//log.Println("Fetching ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body, r.Url), nil
}
