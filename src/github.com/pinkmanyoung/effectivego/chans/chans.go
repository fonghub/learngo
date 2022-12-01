package chans

import (
	"fmt"
	"log"
	"time"
)

var ch = make(chan int)
var remain = 5

func seckill() {
	//使用死循环+阻塞的方式取出数值
	for {
		num := <-ch
		if remain < 1 {
			log.Println("售罄了")
		} else {
			time.Sleep(time.Nanosecond)
			remain = remain - num
			log.Println("购买成功")
		}
	}
}

// 使用channel实现秒杀
func seckillwithchan() {
	//接收者
	go seckill()
	for i := 0; i < 15; i++ {
		ch <- 1
	}
	close(ch)
	log.Println("关闭channel，活动结束")
}

var ch2 = make(chan int)

// 消息通信--客户端
func clientSelect() {
	for {
		num := <-ch
		if remain < 1 {
			log.Println("售罄了")
			ch2 <- 1
			return
		} else {
			time.Sleep(time.Millisecond)
			remain = remain - num
			log.Println("购买成功")
		}
	}
}

// 消息通信--服务端
func serveSelect() {

	go clientSelect()
	for {
		select {
		case _, ok := <-ch2:
			if ok {
				log.Println("接到售罄消息，关闭channel，活动结束")
				close(ch)
				return
			}
		case ch <- 1:
			log.Println("下单成功")
		}
	}

}

// 超时处理机制
func timeout() {
	done := make(chan bool)
	//把阻塞代码放在主程序
	checkoutTimeout(done, 3)
	go func() {
		time.Sleep(time.Second * 5)
		done <- true
	}()
}

// 检测是否超时
func checkoutTimeout(ch chan bool, num time.Duration) {
	select {
	case <-ch:
		fmt.Println("done")
	case <-time.After(time.Second * num):
		fmt.Println("timeout")
	}
}

// 定时任务
func timer() {
	//每秒执行
	t1 := time.Tick(time.Minute)
	//每3秒执行
	t3 := time.Tick(time.Minute * 3)
	//每5秒执行
	t5 := time.Tick(time.Minute * 5)
	for {
		select {
		case <-t1:
			time.Sleep(time.Second)
			log.Println("sec scheduler")
		case <-t3:
			time.Sleep(time.Second)
			log.Println("3 sec scheduler")
		case <-t5:
			time.Sleep(time.Second)
			log.Println("5 sec scheduler")
		}
	}
}

func Call() {
	//使用channel解决并发安全问题
	//seckillwithchan()

	//select用于监听多个channel的读写操作
	//serveSelect()
	//timeout()
	//timer()
}
