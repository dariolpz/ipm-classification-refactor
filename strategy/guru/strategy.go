package guru

import "fmt"

type evictionAlgo interface {
	evict(c *cache)
}

type FIFO struct {
}

func (l *FIFO) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type LRU struct {
}

func (l *LRU) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

type LFU struct {
}

func (l *LFU) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}
