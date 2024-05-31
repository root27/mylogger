package main

import (
	"github.com/root27/mylogger"
	"log"
)

func main() {

	logger, err := mylogger.New("myapp.log", "myapp", mylogger.DEBUG)

	if err != nil {
		log.Fatalf("could not create logger")
	}
	defer logger.Close()

	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	logger.InfoWithColor("this is colorful info")
	logger.WarnWithColor("this is colorful warning")
	logger.ErrorWithColor("this is colorful error")

}
