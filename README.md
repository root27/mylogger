# Logger in Golang

This is customized logging package in Golang. It can be used to log errors, infos, warnings, fatal errors in desired file. 


## Installation

```bash

go get github.com/root27/mylogger

```
---

## Usage

- Without Colors

```go

package main

import (
    "github.com/root27/mylogger"
    "log"
)

func main() {
    logger, err := mylogger.New("app.log", "MYAPP", mylogger.DEBUG)
    
    if err != nil {
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

- With Colors

```go

package main


import (
    "github.com/root27/mylogger"
    "log"
)

func main() {
    logger, err := mylogger.New("app.log", "MYAPP", mylogger.DEBUG)
    
    if err != nil {
		log.Fatalf("could not create logger")
	}

    defer logger.Close()
    
    logger.InfoWithColor("this is colorful info")
    logger.WarnWithColor("this is colorful warning")
    logger.ErrorWithColor("this is colorful error")

}

```


---

## Test Results

```bash

[2024-05-30T13:10:32+03:00] myapp [DEBUG] This is a debug message
[2024-05-30T13:10:32+03:00] myapp [INFO] This is an info message
[2024-05-30T13:10:32+03:00] myapp [WARN] This is a warning message
[2024-05-30T13:10:32+03:00] myapp [ERROR] This is an error message
[2024-05-30T13:10:32+03:00] myapp [FATAL] This is a fatal message

```
- Color Results

$${\color{green}[2024-05-31T17:12:13+03:00] myapp [INFO] this is colorful info}$$<br>
$${\color{red}[2024-05-31T17:12:13+03:00] myapp [ERROR] this is colorful error}$$<br>
$${\color{yellow}[2024-05-31T17:12:13+03:00] myapp [WARN] this is colorful warning}$$

---

## Features

- Colorful Info, Warning and Error messages




