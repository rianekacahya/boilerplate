package main

import (
	"runtime"
	"github.com/rianekacahya/boilerplate/cmd/bootstrap"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	bootstrap.Execute()
}
