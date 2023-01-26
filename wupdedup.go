package main

import (
	"log"
)

func init() {
	initLogger()
	log.Println("init entered")

	log.Println("init exiting..")
}

func main() {
	log.Println("main entered")

	c := loadConfig()
	contexts := loadStorageStrategyContexts(&c)
	for _, context := range contexts {
		context.scanTree()
	}

	log.Println("main exiting..")
}
