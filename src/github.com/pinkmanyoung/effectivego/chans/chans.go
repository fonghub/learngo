package chans

import (
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

func test() {
	go func() {
		_, ok := <-ch
		if ok {
		}
	}()
	ch <- 1
	return
}

func Call() {
	//使用channel解决并发安全问题
	//seckillwithchan()

	//select用于监听多个channel的读写操作
	serveSelect()

	//test()
}
