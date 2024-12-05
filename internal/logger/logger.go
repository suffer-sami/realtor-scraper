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

type ansiColor string

const (
	colorReset  ansiColor = "\x1b[0m"
	colorRed    ansiColor = "\x1b[91m"
	colorGreen  ansiColor = "\x1b[92m"
	colorYellow ansiColor = "\x1b[93m"
	colorPurple ansiColor = "\x1b[95m"
	colorGray   ansiColor = "\x1b[38;5;247m"
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
	l.log(LevelFatal, format, v...)
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
	var colorCode ansiColor

	switch level {
	case LevelDebug:
		levelPrefix = "DEBUG"
		colorCode = colorGray
	case LevelInfo:
		levelPrefix = "INFO"
		colorCode = colorGreen
	case LevelWarn:
		levelPrefix = "WARN"
		colorCode = colorYellow
	case LevelError:
		levelPrefix = "ERROR"
		colorCode = colorRed
	case LevelFatal:
		levelPrefix = "FATAL"
		colorCode = colorPurple
	}

	formattedMsg := fmt.Sprintf("%s%s %-5s%s: %s", l.loggingPrefix, colorCode, levelPrefix, colorReset, format)
	log.Printf(formattedMsg, v...)
}
