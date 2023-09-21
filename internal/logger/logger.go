package logger

import (
	"MovieDataCaptureGo/internal/config"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

var Logger *logrus.Logger

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func setupLogger(cfg *config.Config) *logrus.Logger {
	// init new logger
	Logger := logrus.New()
	var logLevel logrus.Level
	switch cfg.Logger.Level {
	case "debug":
		logLevel = logrus.DebugLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "warn":
		logLevel = logrus.WarnLevel
	default:
		logLevel = logrus.InfoLevel
	}
	// set log level
	Logger.SetLevel(logLevel)

	// set log formatter
	Logger.SetFormatter(&LogFormatter{})

	// set log writers
	var writers []io.Writer
	err := os.MkdirAll(cfg.Logger.LogPath, 0666)
	if err != nil {
		fmt.Printf("cannot create log directory: %v\n", err)
	}
	file, err := os.OpenFile(
		fmt.Sprintf("%s/mdc.log", cfg.Logger.LogPath),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		fmt.Printf("create log file error: %v\n", err)
		writers = append(writers, os.Stdout)
	} else {
		writers = append(writers, os.Stdout)
		writers = append(writers, file)
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	Logger.SetOutput(fileAndStdoutWriter)

	return Logger
}

func init() {
	Logger = setupLogger(config.CFG)
}
