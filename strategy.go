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

// type EmptyContextError struct{}
//
// func (e *EmptyContextError) Error() string {
// 	return fmt.Sprintf("Empty Context Specified for Strategy Pattern")
// }

// Utility function to load concrete StorageStrategy instances from config spec.
// The slice returned is a singleton.
func loadStorageStrategyContexts(c *Config) []*StorageStrategyContext {
	var contexts []*StorageStrategyContext
	if c.Local.Specified() && c.Local.Valid() == true {
		contexts = append(contexts,
			[]*StorageStrategyContext{NewStorageStrategyContext(LocalStrategy{c.Local})}...)
	}
	if c.Smugmug.Specified() && c.Smugmug.Valid() == true {
		contexts = append(contexts,
			[]*StorageStrategyContext{NewStorageStrategyContext(SmugmugStrategy{c.Smugmug})}...)
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
