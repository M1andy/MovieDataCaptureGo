package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"

	. "MovieDataCaptureGo/internal/logger"
)

type JavbusCrawler struct {
	domains string
	crawler *colly.Collector
}

func NewJavbusCrawlerFactory(domains string) *JavbusCrawler {
	if domains == "" {
		domains = "www.javbus.com"
	}
	c := javbusCrawlerFactory(domains)
	return &JavbusCrawler{domains: domains, crawler: c}
}

func javbusCrawlerFactory(domains string) *colly.Collector {
	c := crawlerFactory(domains)

	setupJavbusErrorCallbacks(c)
	setupJavbusXmlCallbacks(c)
	setupJavbusHtmlCallbacks(c)
	setupJavbusRequestsCallbacks(c)
	setupJavbusResponseCallbacks(c)
	return c
}

func (f JavbusCrawler) CrawlNumber(number string) error {
	err := f.crawler.Visit(fmt.Sprintf("https://%s/%s", f.domains, number))
	if err != nil {
		return err
	}
	return nil
}

func setupJavbusErrorCallbacks(c *colly.Collector) {

}

func setupJavbusXmlCallbacks(c *colly.Collector) {
	// TODO this on xml function only works if the movie does not have a director subsection
	// such as https://www.javbus.com/OFJE-377
	c.OnXML("/html/body/div[5]/div[1]/div[2]", func(e *colly.XMLElement) {
		info := &JAVInfo{}

		info.Number = e.ChildText("p[1]/span[2]")
		info.ReleaseDate = e.ChildText("p[2]/text()")
		info.VideoLength = e.ChildText("p[3]/text()")
		info.Director = ""
		info.Studio = e.ChildText("p[4]/a")
		info.Label = e.ChildText("p[5]/a")
		info.Series = e.ChildText("p[6]/a")

		info.Genre = func() []string {
			genres := strings.Split(e.ChildText("p[8]"), "\n")
			for i, genre := range genres {
				genres[i] = strings.TrimSpace(genre)
			}
			return genres
		}()

		info.Actors = func() []string {
			crowedActors := strings.Split(e.ChildText("p[10]"), "\n")
			var actors []string
			for _, actor := range crowedActors {
				tmpActor := strings.TrimSpace(actor)
				if tmpActor != "" {
					actors = append(actors, tmpActor)
				}
			}
			return actors
		}()
		Logger.Infoln(info)
	})
}

func setupJavbusHtmlCallbacks(c *colly.Collector) {
}

func setupJavbusRequestsCallbacks(c *colly.Collector) {
	// TODO read cookie from config or something.
	cookie := ""
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", cookie)
	})
}

func setupJavbusResponseCallbacks(c *colly.Collector) {

}
