package main

import (
	"log"
	"runtime"

	"github.com/ahmdrz/sandogh/src/fileserver"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)
}

func main() {
	svr, err := fileserver.Initialize()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(svr.Run())
}
