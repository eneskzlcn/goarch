package cache

import "time"

/*Cache provides in-memory caching. You can
set the cache content with the expiration duration
which ensures that your data put in to the cache
will be removed after the expiration duration completes.
*/
type Cache interface {
	SetWithExpire(key interface{}, value interface{}, expiration time.Duration) error
	Get(key interface{}) (interface{}, error)
}
