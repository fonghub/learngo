package goroutines

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
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

func waitGp() {
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

var remain int64

// 不使用锁秒杀
func seckill() {
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
	time.Sleep(time.Microsecond)
	remain--
	log.Println("购买成功")
}

// 使用原子操作秒杀
func seckillwithatomic() {
	defer wg.Done()
	res := atomic.AddInt64(&remain, -1)
	if res < 0 {
		log.Println("已售罄")
		return
	}
	time.Sleep(time.Microsecond)
	log.Println("购买成功")
}

// 并发环境下，加锁、不加锁和原子操作的区别
func callseckill() {
	remain = 5
	num := 10
	for i := 0; i < num; i++ {
		wg.Add(1)
		//go seckill()
		//go seckillwithlock()
		go seckillwithatomic()
	}
	wg.Wait()
	log.Println(remain)
}

// 协程中的map会出现并发读错误和并发写错误
func maps() {
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

// 使用sync.map实现map的并发读和并发写
func syncmaps() {
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

type vals struct {
	x int
	y int
}

// 加减法生成工具
func makeHomeWrok(num int, max int) {
	var sm sync.Map
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			x := rand.Intn(max)
			y := rand.Intn(max)
			res := x + y
			sm.Store(vals{x, y}, res)
			wg.Done()
		}()
	}
	wg.Wait()
	var ii = 1
	sm.Range(func(key, value any) bool {
		if ii%4 != 0 {
			fmt.Printf("%v+%v=\t\t", key.(vals).x, key.(vals).y)
		} else {
			fmt.Printf("%v+%v=\n", key.(vals).x, key.(vals).y)
		}
		ii++
		return true
	})
	fmt.Println("-------------")
	ii = 1
	sm.Range(func(key, value any) bool {
		if ii%4 != 0 {
			fmt.Printf("%v-%v=\t\t", key.(vals).x+key.(vals).y, key.(vals).y)
		} else {
			fmt.Printf("%v-%v=\n", key.(vals).x+key.(vals).y, key.(vals).y)
		}
		ii++
		return true
	})
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
	//waitGp()

	//并发不安全，会出现脏读数据；使用锁可以保证并发安全
	//callseckill()

	//协程中的map并发不安全
	//maps()
	//syncmaps()
	makeHomeWrok(100, 20)
}
