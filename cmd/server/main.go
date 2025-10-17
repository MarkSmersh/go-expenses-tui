package main

import (
	"log"
	"log/slog"

	"github.com/MarkSmersh/go-expenses-tui.git/api"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	api.Init()
}
