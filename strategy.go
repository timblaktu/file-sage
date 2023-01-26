package main

import (
	"log"
)

// Strategy-pattern interface for common operations impl by storage providers
type StorageStrategy interface {
	scanTree(c *StorageStrategyContext)
}

// Context used to abstract and "call" different impls at runtime
type StorageStrategyContext struct {
	storageStrategy StorageStrategy
}

func NewStorageStrategyContext(s StorageStrategy) *StorageStrategyContext {
	return &StorageStrategyContext{
		storageStrategy: s,
	}
}

func (c *StorageStrategyContext) setStorageStrategy(s *StorageStrategy) {
	c.storageStrategy = *s
}

func (c *StorageStrategyContext) scanTree() {
	c.storageStrategy.scanTree(c)
}

// Concrete type that implements StorageStrategy interface for Local Storage
type Local struct {
	conf LocalConfig
}

func (s Local) scanTree(c *StorageStrategyContext) {
	log.Println("scanning tree with Local strategy..")
}

// Concrete type that implements StorageStrategy interface for SmugMug
type Smugmug struct {
	conf SmugMugConfig
}

func (s Smugmug) scanTree(c *StorageStrategyContext) {
	log.Println("scanning tree with Smugmug strategy..")
}

// Utility function to load concrete StorageStrategy instances from config spec
func loadStorageStrategyContexts(c *Config) []*StorageStrategyContext {
	var contexts []*StorageStrategyContext
	if c.Local.Specified() {
		contexts = append(contexts,
			[]*StorageStrategyContext{NewStorageStrategyContext(Local{c.Local})}...)
	}
	if c.Smugmug.Specified() {
		contexts = append(contexts,
			[]*StorageStrategyContext{NewStorageStrategyContext(Smugmug{c.Smugmug})}...)
	}
	if len(contexts) == 0 {
		log.Fatalf("No Storage Strategies specified in config")
	} else {
		log.Printf("Loaded %d Storage Strategy Contexts:\n", len(contexts))
		for i, context := range contexts {
			log.Printf("  %d: %+v\n", i, *context)
		}
	}
	return contexts
}
