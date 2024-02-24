package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"

	. "MovieDataCaptureGo/internal/logger"
)

type JavdbCrawler struct {
	domains string
	crawler *colly.Collector
}

func NewJavdbCrawlerFactory(domains string) *JavdbCrawler {
	if domains == "" {
		domains = "(https://|)(www.|)javdb[0-9]*.com(/.*|)"
	}
	allowedDomains := []string{
		domains,
	}
	c := javdbCrawlerFactory(allowedDomains...)
	return &JavdbCrawler{domains: domains, crawler: c}
}

func javdbCrawlerFactory(domains ...string) *colly.Collector {
	c := crawlerFactory(domains...)

	setupJavdbErrorCallbacks(c)
	setupJavdbXmlCallbacks(c)
	setupJavdbHtmlCallbacks(c)
	setupJavdbRequestsCallbacks(c)
	setupJavdbResponseCallbacks(c)
	return c
}

func (f JavdbCrawler) CrawlNumber(number string) error {
	number = strings.ToUpper(number)
	searchUrl := fmt.Sprintf("https://%s/search?q=%s&f=all", f.domains, number)
	err := f.crawler.Visit(searchUrl)
	if err != nil {
		return err
	}
	return nil
}

func setupJavdbErrorCallbacks(c *colly.Collector) {

}

func setupJavdbXmlCallbacks(c *colly.Collector) {

	c.OnXML("/html/body/section/div/div[6]/div[1]", func(e *colly.XMLElement) {
		videoUrl := e.ChildAttr("a", "href")
		Logger.Infoln(videoUrl)
	})
}

func setupJavdbHtmlCallbacks(c *colly.Collector) {
}

func setupJavdbRequestsCallbacks(c *colly.Collector) {
	// TODO read cookie from config or something.
	cookie := ""
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", cookie)
	})
}

func setupJavdbResponseCallbacks(c *colly.Collector) {
	c.OnResponse(func(r *colly.Response) {
		stringCoverter := strings.Builder{}
		stringCoverter.Write(r.Body)
		resp := stringCoverter.String()
		Logger.Debugln(resp)
	})
}
