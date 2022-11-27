package types

import "fmt"

// TestType 类型断言，提供访问接口值底层具体值的方式
func TestType() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // 正确写法，返回一个元组格式 value,type
	fmt.Println(f, ok)

	//接口i没有保存 float64 类型的值
	//以下写法会引发恐慌
	//f = i.(float64)
	//fmt.Println(f)
}

func do(i interface{}) {
	//类型选择
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func Testdo() {
	do(21)
	do("hello")
	do(true)
}
