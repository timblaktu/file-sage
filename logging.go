package main

import "log"

func initLogger() {
	// log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile | log.Lmsgprefix)
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lmsgprefix)
	// log.SetPrefix(": ")
}
