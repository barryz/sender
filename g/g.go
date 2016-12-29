package g

import (
	"log"
	"runtime"
)

const (
	VERSION = "0.0.1@barryz"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
