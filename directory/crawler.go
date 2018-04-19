package directory

import (
	"github.com/sckelemen/crawler/core"
)

type DirectoryCrawler struct {
}

func (dc DirectoryCrawler) CrawlerType() core.CrawlerType {
	return DirectoryCrawlerType
}
func (dc DirectoryCrawler) Crawl(root string) interface{} {
	return dc
}

const (
	DirectoryCrawlerType core.CrawlerType = "directory_crawler"
)
