package glob

import (
	cfg "MovieDataCaptureGo/internal/config"
	. "MovieDataCaptureGo/internal/logger"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var supportVideoExt = []string{
	".mp4",
	".mkv",
	".avi",
	".rmdb",
	".rm",
	".flv",
}

var JAVVideoList []string
var FilesList []string

func globFunc(p string, info os.FileInfo, err error) error {
	if err != nil {
		Logger.Debugln(err)
		return err
	}
	if info.IsDir() {
		return nil
	}
	FilesList = append(FilesList, p)
	return nil
}

func VideoFiles(c *cfg.Config) {
	srcDir := c.Main.SourceDirectory
	err := filepath.Walk(srcDir, globFunc)
	if err != nil {
		Logger.Debugln(err)
	}
	err = filterVideo()
	if err != nil {
		Logger.Debugln(err)
	}
	Logger.Infof("Found %d videos under %s", len(JAVVideoList), srcDir)
}

func filterVideo() error {
	for _, p := range FilesList {
		ext := path.Ext(p)
		if !extOK(ext) {
			continue
		}

		baseName := path.Base(p)
		if baseNameOK(baseName) {
			JAVVideoList = append(JAVVideoList, p)
			Logger.Debugln("Found video: ", p)
		}
	}
	return nil
}

func extOK(ext string) bool {
	for _, s := range supportVideoExt {
		if ext == s {
			return true
		}
	}
	return false
}

func baseNameOK(baseName string) (ok bool) {
	// may add more check method
	return checkBaseNameUsingRegexp(baseName)
}

func checkBaseNameUsingRegexp(baseName string) (ok bool) {
	javMatch, err := regexp.MatchString("[A-Za-z]+-[0-9]+", baseName)
	if err != nil {
		Logger.Debugln(err)
		return false
	}
	fc2Match, err := regexp.MatchString("FC2+-[0-9]+", baseName)
	if err != nil {
		Logger.Debugln(err)
		return false
	}
	return javMatch || fc2Match
}
