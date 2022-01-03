package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println(i)
	i = 1
	return
}
func main() {
	// defer fmt.Println("world")
	// fmt.Print("hello ")

	// for i:=0;i<5;i++ {
	// 	defer fmt.Printf("%d,",i)
	// }
	a()
}

