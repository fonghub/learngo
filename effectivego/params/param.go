package params

import "fmt"

func Into() {
	i := []int{1, 2, 3, 4, 5}
	Param1(i...)
}

func Param1(name ...int) {
	fmt.Printf("name= %d,type is %T", name, name)
	fmt.Println("")
	Param2(name...)
}

func Param2(name ...int) {
	fmt.Println("name", name)
	for k, v := range name {
		fmt.Printf("k=%d,v=%d", k, v)
	}
}
