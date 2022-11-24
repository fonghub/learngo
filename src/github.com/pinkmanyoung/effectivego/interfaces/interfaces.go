package interfaces

import (
	"fmt"
	"math"
)

type abser interface {
	abs() float64
}
type myFloat float64

type vertex struct {
	X, Y float64
}

// myFloat 接收者实现了abs方法
func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// vertex 接收者实现了abs方法
func (v *vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Test1() {
	var a abser
	f := myFloat(-math.Sqrt2)
	v := vertex{3, 4}
	//接口类型的变量可以保存任何实现了这些方法的值
	a = f
	fmt.Println(f.abs())
	// *vertex 指针类型实现了 abs方法，而vertex 值类型没有实现，所以这里只能用指针类型
	a = &v
	fmt.Println(a.abs())
}

// 中学生结构体 middle school student
type mss struct {
	name  string
	class string
}

func (m mss) eat() string {
	str := fmt.Sprintf("%v is eating.", m.name)
	return str
}
func (m mss) sleep() string {
	str := fmt.Sprintf("%v is sleeping.", m.name)
	return str
}

func Test2() {
	var s student
	ms3 := mss{"Tom", "3"}
	// mss 类型实现了 student 定义的所有方法，并且保存了 mss 的值，隐式实现 student 接口
	s = ms3
	fmt.Println(s.eat())
	fmt.Println(s.sleep())
}

func describe(s student) {
	fmt.Printf("(%v, %T)\n", s, s)
}
func Test3() {
	var s student
	ms3 := mss{"Tom", "3"}
	s = ms3
	describe(s)
	//打印：({Tom 3}, interfaces.mss)
	//接口也是值，也具有类型，可以用作函数的参数和返回值
}

func Test4() {
	var s student
	//接口值为nil，方法仍然会被调用
	var ms3 mss
	s = ms3
	describe(s)
	s.eat()
}

func Test5() {
	var s student
	//nil接口值既不保存值也不保存类型，调用方法会产生运行时错误
	describe(s)
	s.eat()
}

func desc(e emptyI) {
	fmt.Printf("(%v, %T)\n", e, e)
}
func Test6() {
	var i emptyI
	desc(i)
	//空接口可以保存任何类型的值
	i = 42
	desc(i)
	i = "hello"
	desc(i)
}
