package internal

import "fmt"

func Q() {
	var i = []int{2, 5, 6, 3, 9, 0, 7, 1}
	fmt.Println(i)

	S(i, 0, len(i)-1)
	fmt.Println(i)

}

// S 快速排序
func S(is []int, left, right int) {
	i := left
	j := right

	if left >= right {
		return
	}
	tmp := is[i]
	for i < j {
		for i < j && is[j] >= tmp {
			j--
		}
		is[i] = is[j]
		for i < j && is[i] <= tmp {
			i++
		}
		is[j] = is[i]
	}
	is[i] = tmp
	S(is, left, i-1)
	S(is, i+1, right)
}
