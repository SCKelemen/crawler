package core

type ICrawler interface {
	Crawl(root string) interface{}
	CrawlerType() CrawlerType
}

type CrawlerType string
