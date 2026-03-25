package dedup

import (
	"crypto/sha1"
	"time"

	lru "github.com/hashicorp/golang-lru"
)

type Dedup struct {
	cache *lru.Cache
	ttl   time.Duration
}

func New(size int, ttl time.Duration) *Dedup {
	c, _ := lru.New(size)
	return &Dedup{cache: c, ttl: ttl}
}

func (d *Dedup) Exists(line []byte) bool {

	key := sha1.Sum(line)
	k := string(key[:])

	if val, ok := d.cache.Get(k); ok {
		if time.Since(val.(time.Time)) < d.ttl {
			return true
		}
	}

	d.cache.Add(k, time.Now())
	return false
}
