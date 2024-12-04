package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type Logger interface {
	Fatalf(string, ...interface{})
	Errorf(string, ...interface{})
	Warnf(string, ...interface{})
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
}

type stdDebugLogger struct {
	mu            sync.Mutex
	level         LogLevel
	loggingPrefix string
}

func NewLogger(prefix string, logLevel string) Logger {

	logLevel = strings.ToUpper(strings.TrimSpace(logLevel))
	var level LogLevel

	switch logLevel {
	case "DEBUG":
		level = LevelDebug
	case "INFO":
		level = LevelInfo
	case "WARN":
		level = LevelWarn
	case "ERROR":
		level = LevelError
	case "FATAL":
		level = LevelFatal
	default:
		level = LevelInfo
	}

	return &stdDebugLogger{
		level:         level,
		loggingPrefix: prefix,
	}
}

func (l *stdDebugLogger) Fatalf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	formattedMsg := fmt.Sprintf("%s FATAL: %s", l.loggingPrefix, format)
	log.Printf(formattedMsg, v...)
	os.Exit(1)
}

func (l *stdDebugLogger) Errorf(format string, v ...interface{}) {
	l.log(LevelError, format, v...)
}

func (l *stdDebugLogger) Warnf(format string, v ...interface{}) {
	l.log(LevelWarn, format, v...)
}

func (l *stdDebugLogger) Infof(format string, v ...interface{}) {
	l.log(LevelInfo, format, v...)
}

func (l *stdDebugLogger) Debugf(format string, v ...interface{}) {
	l.log(LevelDebug, format, v...)
}

func (l *stdDebugLogger) log(level LogLevel, format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level < l.level {
		return
	}

	var levelPrefix string
	switch level {
	case LevelDebug:
		levelPrefix = "DEBUG"
	case LevelInfo:
		levelPrefix = "INFO"
	case LevelWarn:
		levelPrefix = "WARN"
	case LevelError:
		levelPrefix = "ERROR"
	}

	formattedMsg := fmt.Sprintf("%s %s: %s", l.loggingPrefix, levelPrefix, format)
	log.Printf(formattedMsg, v...)
}
