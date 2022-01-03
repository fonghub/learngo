package main

import "fmt"

//空接口
type I interface {}

func main() {
	var i I
	describe(i)

	//空接口可以保存任何类型的值
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n",i,i)
}