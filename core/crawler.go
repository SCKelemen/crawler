package core

// NewCrawler returns a new Crawler (container)
func NewCrawler() Crawler {
	return Crawler{}
}

// Crawler is a type that holds other Crawlers
// it is not a base class
type Crawler struct {
	crawlers []ICrawler
}

// AddCrawler adds an ICrawler to the Crawler container class
func (c Crawler) AddCrawler(crawler ICrawler) bool {
	c.crawlers = append(c.crawlers, crawler)
	return true
}
