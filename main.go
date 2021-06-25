package main

import (
	"runtime"

	"github.org/kbank/app"
	"github.org/kbank/logger"
)

func init() {

	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

func main() {
	logger.Info("Starting the applications ...")
	app.Start()
}
