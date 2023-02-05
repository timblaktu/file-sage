package main

import (
	"log"

	"github.com/timblaktu/wupdedup/config"
	"github.com/timblaktu/wupdedup/db"
)

// -----------------------------------------------------------------------------
// Strategy-pattern interface impl by storage providers
type StorageStrategy interface {
	scanTree(c *StorageStrategyContext)
}

// -----------------------------------------------------------------------------
// Context encapsulates a concrete strategy and enables calling impls at runtime
type StorageStrategyContext struct {
	storageStrategy StorageStrategy
	name            string
	bucket          db.Bucket
	fileCount       int
	nodeCount       int
}

func NewStorageStrategyContext(s StorageStrategy, n string) *StorageStrategyContext {
	return &StorageStrategyContext{
		storageStrategy: s,
		name:            n,
	}
}

func (c *StorageStrategyContext) setStorageStrategy(s *StorageStrategy) {
	c.storageStrategy = *s
}

func (c *StorageStrategyContext) SetBucket(b *db.Bucket) {
	c.bucket = *b
}

func (c *StorageStrategyContext) scanTree() {
	c.storageStrategy.scanTree(c)
}

// -----------------------------------------------------------------------------
// Utility function to load concrete StorageStrategy instances from config spec.
// The slice returned is a singleton.
func loadStorageStrategyContexts(c *config.Config) []*StorageStrategyContext {
	var contexts []*StorageStrategyContext
	if c.Local.Specified() && c.Local.Valid() == true {
		contexts = append(contexts,
			[]*StorageStrategyContext{
				NewStorageStrategyContext(LocalStrategy{c.Local}, "local"),
			}...)
	}
	if c.Smugmug.Specified() && c.Smugmug.Valid() == true {
		contexts = append(contexts,
			[]*StorageStrategyContext{
				NewStorageStrategyContext(SmugmugStrategy{c.Smugmug}, "smugmug"),
			}...)
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
