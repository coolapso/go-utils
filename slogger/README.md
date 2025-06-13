# Slogger

Simple logger built with stdlib log/slog, can be configured with different log levels and log formats

## Usage: 

```go

package main

import (
    "github.com/coolapso/go-utils/slogger"
)

func main.go() {
    logger, _ := slogger.NewLogger("debug", "json")

    logger.Error("This is an error message", "err", err)
}
```
