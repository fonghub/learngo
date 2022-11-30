package goroutines

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func test1(s string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v->%v\n", s, i)
		time.Sleep(time.Millisecond)
	}
}

func test2(str string) {
	for j := 0; j < 5; j++ {
		go log.Printf("%v->%v", str, j)
	}
}

var wg sync.WaitGroup

func waitG() {
	for i := 0; i < 10; i++ {
		//增加协程数
		wg.Add(1)
		go log.Println(i)
		//减少协程数
		wg.Done()
	}
	//等待，直到协程数为0
	wg.Wait()
}

var remain = 5

// 不使用锁秒杀
func seckill() {
	defer wg.Done()
	if remain < 1 {
		log.Println("已售罄")
		return
	}
	time.Sleep(time.Nanosecond)
	remain--
	log.Println("购买成功")
}

var mu sync.Mutex

// 使用锁秒杀
func seckillwithlock() {
	defer wg.Done()
	//加锁
	mu.Lock()
	//解锁
	defer mu.Unlock()
	if remain < 1 {
		log.Println("已售罄")
		return
	}
	time.Sleep(time.Nanosecond)
	remain--
	log.Println("购买成功")
}

func CallSay() {
	//协程序，协程是一种用户态的轻量级线程
	//go test1("world")
	//主程序，主程序结束时不会等待协程执行完毕
	//test1("hello")

	//协程的调度不保证执行的顺序
	//go log.Println("goroutine start")
	//test2("goroutine")
	//go func() { log.Println("使用协程调用匿名函数") }()
	//go log.Println("goroutine end")

	//使用waitgroup可以保证协程全部执行
	//waitG()
	//time.Sleep(time.Second)

	//并发不安全，会出现脏读数据
	//使用锁可以保证并发安全
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	//go seckill()
	//	go seckillwithlock()
	//}
	//wg.Wait()
	//log.Println(remain)

}
