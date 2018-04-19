package web

import (
	"github.com/sckelemen/crawler/core"
)

type WebCrawler struct {
}

func (wc WebCrawler) CrawlerType() core.CrawlerType {
	return WebCrawlerType
}

func (wc WebCrawler) Crawl(root string) interface{} {
	return wc
}

const (
	WebCrawlerType core.CrawlerType = "web_crawler"
)

func NewCrawler() core.ICrawler {
	return WebCrawler{}
}
