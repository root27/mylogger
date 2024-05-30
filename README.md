# Logger in Golang

This is customized logging package in Golang. It can be used to log errors, infos, warnings, fatal errors in desired file. 


## Installation

```bash

go get github.com/root27/mylogger

```

## Usage

```go

package main

import (
    "github.com/root27/mylogger"
    "log"
)

func main() {
    logger := mylogger.New("app.log", "MYAPP", mylogger.DEBUG)
    if logger == nil {
        log.Fatalf("could not create logger")
    }
    defer logger.Close()

    logger.Debug("This is a debug message")
    logger.Info("This is an info message")
    logger.Warn("This is a warning message")
    logger.Error("This is an error message")
    logger.Fatal("This is a fatal message")
}

```
