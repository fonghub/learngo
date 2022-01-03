package main

import (
	"fmt"
	"math"
)

type MyFloat float64

//为类型的接收者声明方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}