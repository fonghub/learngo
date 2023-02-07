package methods

import (
	"fmt"
	"math"
)

// vertex 结构体
type vertex struct {
	X, Y float64
}

// abs 的接收者为v
func (v vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Callabs() {
	v := vertex{3, 4}
	fmt.Println(v.abs())
}

// abs 可以改写成 absFunction
// 方法是一个带接收者参数的函数
func absFunction(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func CallabsFunction() {
	v := vertex{3, 4}
	fmt.Println(absFunction(v))
}

// myFloat 声明一个 float64 类型
type myFloat float64

func (f myFloat) absFloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func CallabsFloat(i int) {
	f := myFloat(i)
	fmt.Println(f.absFloat())
}

// scale 的接收者为 vertesx 类型的指针v
// 指针接收者的方法可以修改接收者指向的值
func (v *vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func Callscale(f float64) {
	v := vertex{3, 4}
	v.scale(f)
	fmt.Println(v.abs())
}

// CallabsBoth 函数使用两种方式调用 abs 方法
// abs 方法是一个 vertex 类型值作为接收者的方法
func CallabsBoth() {
	//传值调用
	v := vertex{3, 4}
	fmt.Println(v.abs())
	//传指针调用
	vv := &v
	fmt.Println(vv.abs())
}

// CallscaleBoth 函数使用两种方式调用 scale 方法
// scale 方法是一个 *vertex 指针类型作为接收者的方法
func CallscaleBoth(f float64) {
	//传值调用
	v := vertex{3, 4}
	v.scale(f)
	fmt.Println(v.abs())
	//传指针调用
	vv := &v
	vv.scale(f)
	fmt.Println(vv.abs())
}

// abs 可以改写成 absFunction
// 方法时一个带接收者参数的函数
func absFunctionpoint(v *vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func CallabsFunctionBoth() {
	v := vertex{15, 20}
	fmt.Println(absFunction(v))
	fmt.Println(absFunctionpoint(&v))
}

//总结：
//1. 接收者无论是指针类型还是值类型的方法，均能接收指针类型或者值类型的调用
//2. 参数为指针类型的方法，只能传入指针为参数；参数为值类型的方法，只能传入值为参数

//选择值或者指针作为接收者？
// 应该选择指针作为接收者
//因为：
//1. 方法能够改变接收者指向的值
//2. 传参时避免复制参数，特别时参数为大型结构体的时候，会更加高效
