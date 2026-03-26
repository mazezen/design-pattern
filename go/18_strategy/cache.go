package strategy

type Cache struct {
	data     map[string]string
	capacity int
	strategy EvictionStrategy

	order []string // 简化实现（记录顺序）
}

func NewCache(cap int, strategy EvictionStrategy) *Cache {
	return &Cache{
		data:     make(map[string]string),
		capacity: cap,
		strategy: strategy,
	}
}

func (c *Cache) SetStrategy(s EvictionStrategy) {
	c.strategy = s
}

func (c *Cache) Get(key string) (string, bool) {
	val, ok := c.data[key]
	if ok {
		c.touch(key)
	}
	return val, ok
}

func (c *Cache) Put(key, value string) {
	if _, ok := c.data[key]; !ok && len(c.data) >= c.capacity {
		c.strategy.Evict(c)
	}

	c.data[key] = value
	c.touch(key)
}

func (c *Cache) touch(key string) {
	// 移除旧位置
	for i, k := range c.order {
		if k == key {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
	// 加到尾部
	c.order = append(c.order, key)
}