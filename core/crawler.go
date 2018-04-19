package core

func NewCrawler() Crawler {
	return Crawler{}
}

type Crawler struct {
	crawlers []ICrawler
}

func (c Crawler) AddCrawler(crawler ICrawler) bool {
	c.crawlers = append(c.crawlers, crawler)
	return true
}
