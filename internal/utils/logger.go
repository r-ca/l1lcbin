package utils

import (
	"fmt"

	// "github.com/samber/do"
)

type Logger interface {
	Debg(msg string)
	Succ(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)

	GetChild(name string) Logger
}

type loggerImpl struct {
	tags []string
}

func (l *loggerImpl) logf(level string, tags []string, msg string) {
	// [DEBG] [main] Hello, world!
	// [INFO] [----] Hello, world!
	// [WARN] [main/sub] Hello, world!

	fmt.Printf("[%s] [%s] %s\n", level, formatTags(tags), msg)
}

func formatTags(tags []string) any {
	if len(tags) == 0 {
		return "----"
	} else if len(tags) == 1 {
		return tags[0]
	} else {
		// join tags with / (all item)
		s := ""
		for i, tag := range tags {
			if i != 0 {
				s += "/"
			}
			s += tag
		}

		return s
	}
}

func NewLogger(tags ...string) Logger {
	return &loggerImpl{
		tags: tags,
	}
}

func (l *loggerImpl) Debg(msg string) {
	l.logf("DEBG", l.tags, msg)
}
func (l *loggerImpl) Succ(msg string) {
	l.logf("SUCC", l.tags, msg)
}
func (l *loggerImpl) Info(msg string) {
	l.logf("INFO", l.tags, msg)
}
func (l *loggerImpl) Warn(msg string) {
	l.logf("WARN", l.tags, msg)
}
func (l *loggerImpl) Error(msg string) {
	l.logf("ERRO", l.tags, msg)
}

func (l *loggerImpl) GetChild(name string) Logger {
	newTags := make([]string, len(l.tags)+1)
	copy(newTags, l.tags)
	newTags[len(l.tags)] = name
	return &loggerImpl{
		tags: newTags,
	}
}

type LoggerFactory interface {
	CreateLogger(tags ...string) Logger
}

type loggerFactoryImpl struct {
}

func (l *loggerFactoryImpl) Create(tags ...string) LoggerFactory {
	return &loggerFactoryImpl{}
}

func (l *loggerFactoryImpl) CreateLogger(tags ...string) Logger {
	return NewLogger(tags...)
}
