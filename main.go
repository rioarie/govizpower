package main

import (
	"govizpower/router"
	"log"
	"runtime"
)

// setRuntime for app
func setRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.Printf("Running With %d CPU\n", nuCPU)
}

func main() {
	setRuntime()
	router.Run()
}
