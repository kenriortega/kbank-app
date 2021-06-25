package main

import (
	"github.org/kbank/app"
	"github.org/kbank/logger"
)

func main() {
	logger.Info("Starting the applications ...")
	app.Start()
}
