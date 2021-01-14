package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	DEBUG      = "DEBUG"
	INFO       = "INFO"
	WARN       = "WARN"
	ERROR      = "ERROR"
	DATEFORMAT = "2006-01-02 15:04:05"
)

var logger *Log

type Log struct {
	file   *os.File
	prefix string
	mu     sync.Mutex
}

func (c *Log) SetPrefix(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.prefix = s
}

func (c *Log) Output(s string) {
	var file string
	var line int
	c.mu.Lock()
	defer c.mu.Unlock()
	var ok bool
	_, file, line, ok = runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	buf := fmt.Sprintf(`[%s] [%s]: %s:%d %s`, time.Now().Format(DATEFORMAT),
		c.prefix, file, line, s,
	)
	os.Stdout.WriteString(buf)
	c.file.WriteString(buf)
}

func InitLogger(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger = &Log{file: file}
	return file
}

func Info(v ...interface{}) {
	logger.SetPrefix(INFO)
	logger.Output(fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	logger.SetPrefix(INFO)
	logger.Output(fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	logger.SetPrefix(WARN)
	logger.Output(fmt.Sprintln(v...))
}

func Warnf(format string, v ...interface{}) {
	logger.SetPrefix(WARN)
	logger.Output(fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	logger.SetPrefix(DEBUG)
	logger.Output(fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	logger.SetPrefix(DEBUG)
	logger.Output(fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	logger.SetPrefix(ERROR)
	logger.Output(fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	logger.SetPrefix(ERROR)
	logger.Output(fmt.Sprintf(format, v...))
}
