package strategy


type EvictionStrategy interface {
	Evict(c *Cache)
}