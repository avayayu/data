package fetcher

import (
	"net/http"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set"
)

type AnnounceCrawler struct {
	CodeSet    mapset.Set
	CodeSetMap map[string]map[string]string
	httpClient *http.Client
	fetchers   []fetcher
}

func NewAnnounceCrawler() *AnnounceCrawler {

	crawler := AnnounceCrawler{}
	crawler.CodeSetMap = make(map[string]map[string]string)
	crawler.fetchers = []fetcher{}
	return &crawler
}

func (crawler *AnnounceCrawler) Crawl() {

	var lock sync.WaitGroup
	lock.Add(2)

	var result chan map[string]map[string]string

	go crawler.crawlSHTotal(&lock, result)

	go crawler.crawlSZTotal(&lock, result)

	lock.Wait()

}

func (crawler *AnnounceCrawler) Serve() {

	ticker := time.NewTicker(20)
	for {
		select {
		case <-ticker.C:

		}
	}

}
