package cache

import (
	"github.com/bluele/gcache"
	"time"
)

type GCacheAdapter struct {
	cache gcache.Cache
}

func NewGCacheAdapter(size int) Cache {
	cache := gcache.New(size).Build()
	return &GCacheAdapter{cache: cache}
}
func (g *GCacheAdapter) SetWithExpire(key interface{}, value interface{}, expiration time.Duration) error {
	return g.cache.SetWithExpire(key, value, expiration)
}
func (g *GCacheAdapter) Get(key interface{}) (interface{}, error) {
	return g.cache.Get(key)
}
