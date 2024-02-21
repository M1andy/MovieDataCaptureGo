package main

import (
	_ "MovieDataCaptureGo/internal/crawler"
	avCawler "MovieDataCaptureGo/internal/crawler"
	. "MovieDataCaptureGo/internal/logger"
	"fmt"
)

func main() {
	Logger.Infoln("Reading config finished!")
	//glob.VideoFiles(c.CFG)

	number := "OFJE-377"
	crawler := avCawler.NewJavbusCrawlerFactory("www.javbus.com")
	err := crawler.CrawlNumber(number)
	if err != nil {
		Logger.Debugln(err)
	}

	outlineCrawler := avCawler.NewAirwikiCrawler("www.airav.wiki")
	err = outlineCrawler.CrawlNumber(number)
	if err != nil {
		Logger.Debugln(err)
	}
	fmt.Scanln()
}
