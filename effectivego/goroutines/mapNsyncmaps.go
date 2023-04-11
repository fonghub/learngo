package goroutines

import (
	"fmt"
	"sync"
)

// Maps 协程中的map会出现并发读错误和并发写错误
func Maps() {
	fmt.Println("---------------Maps---------------")
	m := make(map[int]int)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			m[i] = i
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}

// SyncMaps 使用sync.map实现map的并发读和并发写
func SyncMaps() {
	fmt.Println("---------------SyncMaps---------------")
	var sm sync.Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			sm.Store(i, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	sm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
