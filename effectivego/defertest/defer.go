package defertest

import "fmt"

func DeferTest() bool {
	defer fmt.Println("World.")
	fmt.Print("Hello ")
	return true
}

func DeferMore() bool {
	fmt.Println("before")
	//defer会把函数压入栈中延迟执行，当外出函数返回时，被延迟执行的函数会按照后进先出的顺序调用
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v ", i)
	}
	fmt.Println("after")
	return true
}
