package main

import (
	"log"

	"github.com/timblaktu/wupdedup/db"
)

//-----------------------------------------------------------------------------
// File-level Strategy-pattern interface impl by storage providers
// type FileStrategy interface {
// 	getHdr(c *FileStrategyContext) ([]byte, error)
// }
//
// // Context used to abstract and "call" different impls at runtime
// type FileStrategyContext struct {
// 	storageStrategy FileStrategy
// 	name            string
// }
//
// func NewFileStrategyContext(s FileStrategy, n string) *FileStrategyContext {
// 	return &FileStrategyContext{
// 		storageStrategy: s,
// 		name:            n,
// 	}
// }
//
// func (c *FileStrategyContext) setFileStrategy(s *FileStrategy) {
// 	c.storageStrategy = *s
// }
//
// func (c *FileStrategyContext) scanTree() {
// 	c.storageStrategy.scanTree(c)
// }
//-----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Root-level Strategy-pattern interface impl by storage providers
type StorageStrategy interface {
	scanTree(c *StorageStrategyContext)
}

// Context used to abstract and "call" different impls at runtime
type StorageStrategyContext struct {
	storageStrategy StorageStrategy
	name            string
	bucket          db.Bucket
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

//-----------------------------------------------------------------------------

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
