package fetcher

import (
	"net/http"
	"sync"
)

func getSZReq(startDate, endDate string, page uint16) (*http.Request, string) {

	return nil, ""
}

func (crawler *AnnounceCrawler) SZ(page uint16) (announce *SHAnnouce, err error) {

	return nil, nil
}

func (crawler *AnnounceCrawler) crawlSZTotal(group *sync.WaitGroup, data chan map[string]map[string]string) {
	defer group.Done()

}
