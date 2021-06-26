package main

import (
	"runtime"

	api "github.org/kbank/cmd/api"
	"github.org/kbank/internal/logger"
)

func init() {

	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

func main() {
	logger.Info("Starting the applications ...")
	api.Start()
}
