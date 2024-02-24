package main

import (
	"fmt"

	. "MovieDataCaptureGo/internal/config"
	_ "MovieDataCaptureGo/internal/crawler"
	avCawler "MovieDataCaptureGo/internal/crawler"
	"MovieDataCaptureGo/internal/glob"
	. "MovieDataCaptureGo/internal/logger"
)

func main() {
	Logger.Infoln("Reading config finished!")
	_, err := glob.JAVFiles(CFG.Main.SourceDirectory)
	if err != nil {
		Logger.Debugln(err)
	}
	number := "IPX-901"
	crawler := avCawler.NewJavdbCrawlerFactory("www.javdb524.com")
	err = crawler.CrawlNumber(number)
	if err != nil {
		Logger.Debugln(err)
	}

	Logger.Infoln("Waiting for exit, press any key...")
	fmt.Scanln()
}
