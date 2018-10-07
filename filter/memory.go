package filter

import (
	"github.com/GerryLon/go-crawler/utils/text"
)

// de duplication using memory, just a map[string][string]
type MemoryDedupFilter struct {
	DuplicateFilter
	store map[string]string
}

// init redis connection pool
func (filter *MemoryDedupFilter) init() {
	filter.store = make(map[string]string)
}

// get one connection from redis connection pool
func (filter *MemoryDedupFilter) getMemoryStore() map[string]string {
	if filter.store == nil {
		filter.init()
	}
	return filter.store
}

// redis: exists key
func (filter *MemoryDedupFilter) Has(key string) bool {
	store := filter.getMemoryStore()
	_, ok := store[text.MD5(key)]
	return ok
}

//redis: get mdt(key)
func (filter *MemoryDedupFilter) Get(key string) string {
	store := filter.getMemoryStore()
	return store[text.MD5(key)]
}

// memory: set key if it dose not exist
func (filter *MemoryDedupFilter) setNX(key string) bool {
	store := filter.getMemoryStore()

	md5OfKey := text.MD5(key)

	_, ok := store[md5OfKey]
	if ok {
		return false
	}

	store[md5OfKey] = key
	return true
}

// set using redis: setnx key
func (filter *MemoryDedupFilter) Set(key string) bool {
	return filter.setNX(key)
}

// delete key from store
func (filter *MemoryDedupFilter) Del(key string) bool {
	store := filter.getMemoryStore()
	delete(store, text.MD5(key))
	return true
}
