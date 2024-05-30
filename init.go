package mylogger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	out    io.Writer
	mu     sync.Mutex
	level  LogLevel
	prefix string
}

func New(filePath, prefix string, level LogLevel) (*Logger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Logger{out: file, prefix: prefix, level: level}, nil
}

func (l *Logger) log(level LogLevel, msg string) {
	if level < l.level {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Fprintf(l.out, "[%s] %s [%s] %s\n", timestamp, l.prefix, level.String(), msg)
}

func (level LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[level]
}

