package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	//显示实现 implements
	// var i I 
	// t := T{"hello"}
	// i = t

	//隐式实现
	var i I = T{"hello"}
	i.M()
}