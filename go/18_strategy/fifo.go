package strategy

type FIFO struct{}

func (f *FIFO) Evict(c *Cache) {
	if len(c.order) == 0 {
		return
	}

	// 先进先出（头部）
	first := c.order[0]

	delete(c.data, first)
	c.order = c.order[1:]
}