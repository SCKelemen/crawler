package web

import (
	"github.com/sckelemen/crawler/core"
)

// WebCrawler is an implementation of ICrawler
// for crawling websites. It traverses links
// using dfs
type WebCrawler struct {
}

// CrawlerType returns the identifier for the
// WebCrawler type
func (wc WebCrawler) CrawlerType() core.CrawlerType {
	return WebCrawlerType
}

// Crawl crawls the root string
func (wc WebCrawler) Crawl(root string) interface{} {
	return wc
}

const (
	// WebCrawlerType is the indentifier for the
	// WebCrawler implementation
	WebCrawlerType core.CrawlerType = "web_crawler"
)

// NewCrawler returns an instance of WebCrawler
func NewCrawler() WebCrawler {
	return WebCrawler{}
}
