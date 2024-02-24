package crawler

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"

	. "MovieDataCaptureGo/internal/logger"
)

type OutlineInfo struct {
	Title   string
	outline string
}
type AirwikiCrawler struct {
	domains string
	crawler *colly.Collector
}

func NewAirwikiCrawler(domains string) *AirwikiCrawler {
	c := airwikiCrawlerFactory(domains)
	return &AirwikiCrawler{domains: domains, crawler: c}
}

func airwikiCrawlerFactory(domains string) *colly.Collector {
	c := crawlerFactory(domains)

	setupAirwikiXmlCallbacks(c)
	return c
}

func (f AirwikiCrawler) CrawlNumber(number string) error {
	url := fmt.Sprintf("https://%s/video/%s", f.domains, number)
	err := f.crawler.Visit(url)
	if err != nil {
		return fmt.Errorf("visit url %s failed: %v", url, err)
	}
	return nil
}

func setupAirwikiXmlCallbacks(c *colly.Collector) {
	c.OnXML("//*[@id=\"__next\"]/div[2]/div[2]/div[2]", func(e *colly.XMLElement) {
		time.Sleep(time.Second * 2)
		info := &OutlineInfo{}
		info.Title = e.ChildText("div[4]/p")
		info.outline = e.ChildText("div[5]/div[2]/div/h5[1]")
		Logger.Infoln(info)
	})
}

type AiravCrawler struct {
	domains string
	crawler *colly.Collector
}

func NewAiravCrawler(domains string) *AirwikiCrawler {
	c := airavCrawlerFactory(domains)
	return &AirwikiCrawler{domains: domains, crawler: c}
}

func airavCrawlerFactory(domains string) *colly.Collector {
	c := crawlerFactory(domains)

	setupAiravXmlCallbacks(c)
	return c
}

func (f AiravCrawler) CrawlNumber(number string) error {
	url := fmt.Sprintf("https://%s/video/%s", f.domains, number)
	err := f.crawler.Visit(url)
	if err != nil {
		return fmt.Errorf("visit url %s failed: %v", url, err)
	}
	return nil
}

func setupAiravXmlCallbacks(c *colly.Collector) {
	c.OnXML("//*[@id=\"__next\"]/div[2]/div[2]/div[2]", func(e *colly.XMLElement) {
		time.Sleep(time.Second * 2)
		info := &OutlineInfo{}
		info.Title = e.ChildText("div[4]/p")
		info.outline = e.ChildText("div[5]/div[2]/div/h5[1]")
		Logger.Infoln(info)
	})
}
