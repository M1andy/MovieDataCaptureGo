package glob

import (
	"os"
	"path"
	"path/filepath"
	"regexp"

	. "MovieDataCaptureGo/internal/logger"
)

var supportedVideoExt = []string{
	".mp4",
	".mkv",
	".avi",
	".rmdb",
	".rm",
	".flv",
}

var filesList []string

func globFunc(p string, info os.FileInfo, err error) error {
	if err != nil {
		Logger.Debugln(err)
		return err
	}
	if info.IsDir() {
		return nil
	}
	ext := path.Ext(p)
	if !extOK(ext) {
		return nil
	}

	baseName := path.Base(p)
	if baseNameOK(baseName) {
		filesList = append(filesList, p)
		Logger.Debugln("Found video: ", p)
	}

	return nil
}

func JAVFiles(srcDir string) ([]string, error) {
	err := filepath.Walk(srcDir, globFunc)
	if err != nil {
		return nil, err
	}
	Logger.Infof("Found %d videos under %s", len(filesList), srcDir)
	return filesList, err
}

func extOK(ext string) bool {
	for _, s := range supportedVideoExt {
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
