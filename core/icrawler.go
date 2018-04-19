package core

type ICrawler interface {
	Crawl(root string) interface{}
	CrawlerType() CrawlerType
}

type CrawlerType string

// the crawlers will be feed input strings representing the root node
// for the web crawler implementation, I will likely use dfs so I don't
// exhaust the host, but for the directory implementation, it's unlikely
// to cause a problem using bfs. I think that would be more efficient

// crawlers will emit their findings
