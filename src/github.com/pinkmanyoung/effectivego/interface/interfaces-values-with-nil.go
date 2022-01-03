package main

import "fmt"

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	//接口值为nil
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n",i,i)
}

func main() {
	var i I

//显式实现接口
	var t *T
	i = t
	describe(i)
	i.M()

//隐式实现接口
	i = &T{"hello"}
	describe(i)
	i.M()
}