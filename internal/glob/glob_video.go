package glob

import (
	"os"
	"path"
	"path/filepath"
	"regexp"

	"MovieDataCaptureGo/internal/crawler"
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

var filesList []crawler.JAVInfo

func globFunc(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// filter directory
	if info.IsDir() {
		return nil
	}

	// check file extension
	fileExt := path.Ext(p)
	if !extOK(fileExt) {
		return nil
	}

	// check file size, must larger than 300MB
	fileSize := info.Size()
	if fileSize <= 314572800 {
		return nil
	}

	// check fileName
	fileName := info.Name()
	fileNameWithoutExt := fileName[0 : len(fileName)-len(fileExt)]
	if !isJavNumber(fileNameWithoutExt) {
		return nil
	}

	// append to fileList
	filesList = append(filesList, crawler.JAVInfo{Number: fileNameWithoutExt, FilePath: p})
	Logger.Debugf("Found: %s | Path: %s \n", fileNameWithoutExt, p)

	return nil
}

func JAVFiles(srcDir string) ([]crawler.JAVInfo, error) {
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

func isJavNumber(baseName string) (ok bool) {
	// may add more check method
	return checkBaseNameUsingRegexp(baseName)
}

func checkBaseNameUsingRegexp(fileName string) (ok bool) {
	javMatch, err := regexp.MatchString("[A-Za-z]+-[0-9]+(-c|-C|)", fileName)
	if err != nil {
		Logger.Debugln(err)
		return false
	}
	fc2Match, err := regexp.MatchString("FC2+(-PPV|)-[0-9]+", fileName)
	if err != nil {
		Logger.Debugln(err)
		return false
	}
	return javMatch || fc2Match
}
