package core

// ICrawler is an interface for different types of
// crawlers. These include directory crawlers, and
// web crawlers, at the moment.
type ICrawler interface {
	Crawl(root string) interface{}
	CrawlerType() CrawlerType
}

// CrawlerType is a string to identity the type of
// crawler, but the real purpose to ensure the
// concrete impelementation only satisfies the
// ICrawler interface
type CrawlerType string

// the crawlers will be feed input strings representing the root node
// for the web crawler implementation, I will likely use dfs so I don't
// exhaust the host, but for the directory implementation, it's unlikely
// to cause a problem using bfs. I think that would be more efficient

// crawlers will emit their findings
