package internal

import (
	"fmt"
	"reflect"
)

// Func00 返回变量的类型名和种类名
func Func00(m map[int]any) {
	for _, v := range m {
		typeOf := reflect.TypeOf(v)
		if typeOf != nil {
			fmt.Println(typeOf.Name(), typeOf.Kind())
		}
	}
}

func Func0() {
	var m = map[int]any{}
	m[0] = 1
	m[1] = "hello"
	m[2] = []int{1, 2, 3}
	m[3] = nil
	Func00(m)
}
