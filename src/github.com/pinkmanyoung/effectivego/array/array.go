package main

import "fmt"

func main() {
	var a [2]string
	a = [2]string{"Hello","world"}
	fmt.Println(a)
	test(&a)
	fmt.Println(a)
}

func test(arr *[2]string){
	arr[1] = "World"
	fmt.Println(*arr)
}