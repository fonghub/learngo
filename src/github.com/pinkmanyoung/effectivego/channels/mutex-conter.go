package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}	

func main() {
	m := make(map[string]int)
	c := SafeCounter{v:m}
	for i:=0;i<1000;i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}