package crawler

import (
	"MovieDataCaptureGo/internal/proxy"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type AVCrawler interface {
	CrawlNumber(number string) error
}

func crawlerFactory(domains string) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(domains),
		colly.Async(true),
	)
	c.SetProxyFunc(proxy.ProxyFunc)
	extensions.RandomUserAgent(c)
	return c
}
