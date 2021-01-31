package fetcher

type fetcher interface { //
	Name() string
	fetch(code string) interface{}
}

func (crawler *AnnounceCrawler) AddFetcher() {}
