package directory

import (
	"os"
	"path/filepath"

	"github.com/sckelemen/crawler/core"
)

func New() DirectoryCrawler {
	file := make(chan string)
	dir := make(chan string)
	return DirectoryCrawler{EmitFile: file, EmitDir: dir}
}

type DirectoryCrawler struct {
	files       []string
	directories []string
	EmitFile    chan string
	EmitDir     chan string
}

func (dc DirectoryCrawler) CrawlerType() core.CrawlerType {
	return DirectoryCrawlerType
}
func (dc DirectoryCrawler) Crawl(root string) interface{} {
	filepath.Walk(root, dc.visit)
	return dc
}

const (
	DirectoryCrawlerType core.CrawlerType = "directory_crawler"
)

func (dc DirectoryCrawler) visit(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		// emit directory
		dc.EmitDir <- path
		// don't continue processing this
		// we wont to process on a different thread
		return filepath.SkipDir
	}
	dc.EmitFile <- path
	return nil
}

func NewCrawler() core.ICrawler {
	return DirectoryCrawler{}
}
