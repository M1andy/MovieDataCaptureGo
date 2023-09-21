package main

import (
	c "MovieDataCaptureGo/internal/config"
	"MovieDataCaptureGo/internal/glob"
	. "MovieDataCaptureGo/internal/logger"
	"fmt"
)

func main() {
	Logger.Infoln("Reading config finished!")
	glob.VideoFiles(c.CFG)

	Logger.Infoln("Waiting for exit, press any key...")
	fmt.Scanln()
}
