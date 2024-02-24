package crawler

import (
	"github.com/corpix/uarand"
	"github.com/gocolly/colly/v2"

	. "MovieDataCaptureGo/internal/logger"
	"MovieDataCaptureGo/internal/proxy"
)

type AVCrawler interface {
	CrawlNumber(number string) error
}

func crawlerFactory(domains ...string) *colly.Collector {
	// init new collector
	c := colly.NewCollector(
		//colly.AllowedDomains(domains...),
		colly.Async(true),
	)
	c.SetProxyFunc(proxy.ProxyFunc)
	setupCommonCallbacks(c)
	return c
}

func setupCommonCallbacks(c *colly.Collector) {
	// random ua
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", uarand.GetRandom())
	})

	// error logging
	c.OnError(func(r *colly.Response, err error) {
		Logger.Debugf("status: %d | %s\n", r.StatusCode, err)
	})
}

type JAVInfo struct {
	Title       string
	Number      string
	ReleaseDate string
	VideoLength string
	Director    string
	Studio      string
	Label       string
	Series      string
	Genre       []string
	Actors      []string
	Outline     string
	Score       float32
}
