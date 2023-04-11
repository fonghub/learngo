package goroutines

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

// WaitGp 使用WaitGroup可以保证协程全部执行
func WaitGp() {
	for i := 0; i < 10; i++ {
		//增加协程数
		wg.Add(1)
		go func(i int) {
			log.Println(i)
			//减少协程数，wg.Done()方法需要在协程函数内执行
			wg.Done()
		}(i)
	}
	//等待，直到协程数为0
	wg.Wait()
}
