package main

import "fmt"

//类型断言 提供访问接口值底层具体值的方式
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s,ok := i.(string)
	fmt.Println(s,ok)

	f,ok := i.(float64) // 正确写法，返回一个元组格式 value,type
	fmt.Println(f,ok)

	//接口i没有保存float64类型的值
	//以下写法会引发恐慌
	f = i.(float64)
	fmt.Println(f)
}