package functions

import (
	"fmt"
	"math"
)

// 定义一个匿名函数
var fn = func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// Res 定义一个函数类型
type res func(x, y float64) float64

// 函数作为参数
func compute(fn res) float64 {
	return fn(3, 4)
}

func callCompute() {
	//直接调用匿名函数
	fmt.Println(fn(5, 12))
	// 函数作为参数
	fmt.Println(compute(fn))
}

// 函数作为返回值
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func callAdder() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			i,
			pos(i),
			neg(-2*i),
		)
	}
}

// 斐波那契数列
func fibonacci() func() int {
	res := 0
	i, j := 0, 1
	return func() int {
		res = i
		i, j = j, i+j
		return res
	}
}

func callfibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", f())
	}
}

// FunctionCall 调用函数
func FunctionCall() {

	//callCompute()
	//callAdder()
	callfibonacci()
}
