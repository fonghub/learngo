package main

import (
	"fmt"
	"math"
)

//函数作为参数
func compute(fn func(float64,float64) float64) float64 {
	return fn(3,4)
}

func main() {
	//匿名函数，函数作为返回值
	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x+y*y)
	}

	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))

	fmt.Println(compute(math.Pow))
}