package main

import (
	"flag"

	"github.com/sckelemen/crawler/core"
	"github.com/sckelemen/crawler/directory"
	"github.com/sckelemen/crawler/web"
)

func main() {
	flag.Parse()
	workerId := flag.Arg(0)
	webcrawler := web.WebCrawler{}
	dircrawler := directory.DirectoryCrawler{}
	spider := Crawler{}
	spider.AddCrawler(webcrawler)
	spider.AddCrawler(dircrawler)

}

type Crawler struct {
	crawlers []core.ICrawler
}

func (c Crawler) AddCrawler(crawler core.ICrawler) bool {
	c.crawlers = append(c.crawlers, crawler)
	return true
}
