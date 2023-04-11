package goroutines

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// 库存
var remain int64

// secKill 不使用锁秒杀，100%超卖
func secKill() {
	defer wg.Done()
	if remain < 1 {
		log.Println("已售罄")
		return
	}
	//暂停1毫秒，模拟程序运行
	//程序执行时间越长，并发越不安全
	time.Sleep(time.Microsecond)
	remain--
	log.Println("购买成功")
}

// 排他锁
var mu sync.Mutex

// secKillWithLock 使用锁秒杀，秒杀正常
func secKillWithLock() {
	defer wg.Done()
	//加锁
	mu.Lock()
	//解锁
	defer mu.Unlock()
	if remain < 1 {
		log.Println("已售罄")
		return
	}
	time.Sleep(time.Microsecond)
	remain--
	log.Println("购买成功")
}

// secKillWithAtomic 使用原子操作秒杀，秒杀正常
func secKillWithAtomic() {
	defer wg.Done()
	res := atomic.AddInt64(&remain, -1)
	if res < 0 {
		log.Println("已售罄")
		return
	}
	time.Sleep(time.Microsecond)
	log.Println("购买成功")
}

// CallSecKill 并发环境下，加锁、不加锁和原子操作的区别
func CallSecKill() {
	remain = 5
	num := 10
	for i := 0; i < num; i++ {
		wg.Add(1)
		//go secKill()
		//go secKillWithLock()
		go secKillWithAtomic()
	}
	wg.Wait()
	log.Println(remain)
}
