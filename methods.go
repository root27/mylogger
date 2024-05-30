package mylogger

import "os"

func (l *Logger) Debug(msg string) {
	l.log(DEBUG, msg)
}

func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

func (l *Logger) Warn(msg string) {
	l.log(WARN, msg)
}

func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}

func (l *Logger) Fatal(msg string) {
	l.log(FATAL, msg)
	os.Exit(1)
}

func (l *Logger) Close() error {
	if file, ok := l.out.(*os.File); ok {
		return file.Close()
	}
	return nil
}
