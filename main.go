package main

import (
	"github.com/sckelemen/crawler/core"
	"github.com/sckelemen/crawler/directory"
	"github.com/sckelemen/crawler/web"
)

func main() {
	webcrawler := web.NewCrawler()
	dircrawler := directory.NewCrawler()
	spider := core.NewCrawler()
	spider.AddCrawler(webcrawler)
	spider.AddCrawler(dircrawler)
}
