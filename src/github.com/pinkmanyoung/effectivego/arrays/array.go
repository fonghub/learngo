package arrays

import "fmt"

var arr [10]int

// 生产10个单数
func Aget() [10]int {
	num := 1
	for i := 0; i < 10; i++ {
		arr[i] = num
		num += 2
	}
	return arr
}

// 打印数组
func Aprint(arr [10]int) {
	fmt.Println("start")
	for k, v := range arr {
		fmt.Printf("k=%d,v=%d\n", k, v)
	}
	fmt.Println("end")
}
