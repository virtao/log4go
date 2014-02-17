package main

import (
	"time"
)

import l4g "github.com/virtao/log4go"

func main() {
	log := make(l4g.Logger, 0)
	log.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Close()
}
