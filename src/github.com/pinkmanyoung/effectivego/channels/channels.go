package main

import "fmt"

//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//
//	c <- sum
//}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 6)
	//go sum(s[:len(s)/2],c)
	//go sum(s[len(s)/2:],c)
	//
	//x,y := <-c,<-c
	//
	//fmt.Println(x,y,x+y)
	for _, v := range s {
		c <- v
	}
	for v := range c {
		fmt.Println(v)
	}
}
