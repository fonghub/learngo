package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}

//指针类型接收者的方法，既能接收值调用，也能接受指针调用
func (v *Vertex) Scale(f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

//带有指针参数的方法，只能传入指针
func ScaleFunc(v *Vertex,f float64) {
	v.X = v.X*f 
	v.Y = v.Y*f 
}

//值类型接收者的方法，既能接收值调用，也能接受指针调用
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

//带有值参数的方法，只能传入值
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main() {
	v := Vertex{3,4}
	v.Scale(2)
	ScaleFunc(&v,10)

	p := &Vertex{4,3}
 	p.Scale(3)
 	ScaleFunc(p,8)

 	fmt.Println(v,p)
//---------------------------//
	v = Vertex{3,4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p = &Vertex{4,3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

}

//总结：
//1. 接收者无论是指针类型还是值类型的方法，均能接收指针类型或者值类型的调用
//2. 参数为指针类型的方法，只能传入指针为参数；参数为值类型的方法，只能传入值为参数

//选择值或者指针作为接收者？
// 应该选择指针作为接收者
//因为：
//1. 方法能够改变接收者指向的值
//2. 传参时避免复制参数，特别时参数为大型结构体的时候，会更加高效