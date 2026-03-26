package strategy

type LRU struct{}

func (l *LRU) Evict(c *Cache) {
	if len(c.order) == 0 {
		return
	}

	// 最久未使用（头部）
	oldest := c.order[0]

	delete(c.data, oldest)
	c.order = c.order[1:]
}