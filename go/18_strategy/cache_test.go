package strategy

import "testing"

func TestLRU(t *testing.T) {
	c := NewCache(2, &LRU{})

	c.Put("a", "1")
	c.Put("b", "2")

	c.Get("a")     // a 变新
	c.Put("c", "3") // 应该淘汰 b

	if _, ok := c.Get("b"); ok {
		t.Errorf("b should be evicted")
	}

	if _, ok := c.Get("a"); !ok {
		t.Errorf("a should exist")
	}
}

func TestFIFO(t *testing.T) {
	c := NewCache(2, &FIFO{})

	c.Put("a", "1")
	c.Put("b", "2")
	c.Put("c", "3") // 淘汰 a

	if _, ok := c.Get("a"); ok {
		t.Errorf("a should be evicted")
	}
}

func TestSwitchStrategy(t *testing.T) {
	c := NewCache(2, &FIFO{})

	c.Put("a", "1")
	c.Put("b", "2")

	c.SetStrategy(&LRU{})

	c.Get("a")
	c.Put("c", "3") // 按 LRU 淘汰 b

	if _, ok := c.Get("b"); ok {
		t.Errorf("b should be evicted")
	}
}