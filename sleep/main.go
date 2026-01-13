package main

import (
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: sleep <duration>")
		os.Exit(1)
	}
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		println("Invalid duration:", err.Error())
		os.Exit(1)
	}
	time.Sleep(duration)
}
