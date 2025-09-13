package main

import (
	"fmt"
	"sync"
	"time"
)

type Safecounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *Safecounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *Safecounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() // defer so we unlock the mutex after return
	return c.v[key]
}

func main() {
	c := Safecounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("randomKey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("randomKey"))
}
