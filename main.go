package main

import (
	"govizpower/router"
	"log"
	"runtime"
)

func setRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.Printf("Running With %d CPU\n", nuCPU)
}

func main() {
	setRuntime()
	router.Run()
}
