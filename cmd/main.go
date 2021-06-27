package main

import (
	"log"
	"runtime"
	"github.com/rianekacahya/boilerplate/cmd/bootstrap"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	runtime.GOMAXPROCS(runtime.NumCPU())

	// bootstrap dependencies
	bootstrap.Dependencies()

	// bootstrap command
	bootstrap.Execute()
}
