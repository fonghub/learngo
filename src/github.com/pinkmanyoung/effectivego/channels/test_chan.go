package main

import (
	"fmt"
	"time"
)

func in(c chan int) {
	for i:=0;i<10;i++ {
		c <- i
		fmt.Printf("in=%d\n",i)
		time.Sleep(10 * time.Millisecond)
	}
	close(c)
}

func out(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}


func main() {
	c := make(chan int,50)
	fmt.Println("start")
	go out(c)
	in(c)
	fmt.Println("end")
}

