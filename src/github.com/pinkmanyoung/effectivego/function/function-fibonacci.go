package main

import "fmt"

func fibonacci() func() int {
	x,y := 0,1
	var temp int
	return func() int {
		temp,x,y = x,y,x+y
		return temp
	}
}

func main() {
	f := fibonacci()
	for i:=0;i<10;i++ {
		fmt.Println(f())
	}
}