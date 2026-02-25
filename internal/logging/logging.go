package logging

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/glebateee/basic/internal/config"
)

type LogLevel uint

const (
	Trace LogLevel = iota
	Debug
	Info
	Warn
	Fatal
	None
)

type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})
	Debug(string)
	Debugf(string, ...interface{})
	Info(string)
	Infof(string, ...interface{})
	Warn(string)
	Warnf(string, ...interface{})
	Panic(string)
	Panicf(string, ...interface{})
}

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

var stringToLevel = map[string]LogLevel{
	"TRACE": Trace,
	"DEBUG": Debug,
	"INFO":  Info,
	"WARN":  Warn,
	"FATAL": Fatal,
}

func New(cfg config.Config) Logger {
	level := cfg.GetStringDefault("logging:level", "DEBUG")
	logLevel, ok := stringToLevel[strings.ToUpper(level)]
	if !ok {
		logLevel = Debug
	}
	flags := log.Lmsgprefix | log.Ltime
	return &DefaultLogger{
		minLevel: logLevel,
		loggers: map[LogLevel]*log.Logger{
			Trace: log.New(os.Stdout, "TRACE ", flags),
			Debug: log.New(os.Stdout, "DEBUG ", flags),
			Info:  log.New(os.Stdout, "INFO ", flags),
			Warn:  log.New(os.Stdout, "WARN ", flags),
			Fatal: log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(msgLevel LogLevel, msg string) {
	if msgLevel >= l.minLevel {
		l.loggers[msgLevel].Println(msg)
	}
}

func (l *DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

func (l *DefaultLogger) Tracef(pattern string, args ...interface{}) {
	l.write(Trace, fmt.Sprintf(pattern, args...))
}

func (l *DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

func (l *DefaultLogger) Debugf(template string, args ...interface{}) {
	l.write(Debug, fmt.Sprintf(template, args...))
}

func (l *DefaultLogger) Info(msg string) {
	l.write(Info, msg)
}

func (l *DefaultLogger) Infof(template string, args ...interface{}) {
	l.write(Info, fmt.Sprintf(template, args...))
}

func (l *DefaultLogger) Warn(msg string) {
	l.write(Warn, msg)
}

func (l *DefaultLogger) Warnf(template string, args ...interface{}) {
	l.write(Warn, fmt.Sprintf(template, args...))
}

func (l *DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

func (l *DefaultLogger) Panicf(template string, args ...interface{}) {
	formattedMsg := fmt.Sprintf(template, args...)
	l.write(Fatal, formattedMsg)
	if l.triggerPanic {
		panic(formattedMsg)
	}
}
