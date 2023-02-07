package arrays

import "fmt"

// 生产10个单数
func Sget(n int) []int {
	//创建切片
	slc := make([]int, n)
	num := 1
	for i := 0; i < n; i++ {
		slc[i] = num
		num += 2
	}
	return slc
}

// 打印切片
func Sprint(s []int) {
	fmt.Println("start")
	for k, v := range s {
		fmt.Printf("k=%d,v=%d\n", k, v)
	}
	fmt.Println("end")
}

// 打印切片和切片的长度和容量
// 切片的长度len()就是切片包含的元素个数
// 切片的容量cap()就是从切片的第一个元素开始，到底层数组元素末尾的个数
func PLenCap(s []int) {
	fmt.Printf("%v len=%d cap=%d \n", s, len(s), cap(s))
}

// 切片只能向右扩展，不能向左扩展
// 切片从左边被截取掉，不能还原，右边则可以
func DoSlice(s []int) {
	PLenCap(s)

	//清空切片
	s = s[:0]
	PLenCap(s)

	//还原切片
	s = s[:cap(s)]
	PLenCap(s)

	//从左边截半数组，数组容量减半
	s = s[cap(s)/2:]
	PLenCap(s)

	//从右边截半数组，数组容量不变
	//s = s[:cap(s)/2]
	//PLenCap(s)

	//还原切片，向右扩展
	s = s[:cap(s)]
	PLenCap(s)

	//从左边截半切片
	//s = s[len(s)/2:]
	//PLenCap(s)
}

// 向切片追加元素
func Append(s []int, e int) {
	s = append(s, e)
	PLenCap(s)
}

// 使用append()函数
// 返回一个包含n个元素的切片
func AppendSlice(n int) []int {
	var slc []int
	for i := 0; i < n; i++ {
		//使用内建函数append向切片内追加元素
		//当切片的len大于cap，切片会改变cap的大小
		//本质是重新分配内存，把原来的数据拷贝过来
		//如果每一次都是把cap改为和len一样大，结果是每一次调用append函数，都要重新分配内存再做数据拷贝，效率不大
		//所以当len大于2^2时，分配给切片的容量cap=2*2^2
		//所以当len=5时，cap=8
		//所以当len=8时，cap=16
		//所以当len=2^n时，cap=2*2^n
		slc = append(slc, i)
	}
	return slc
}
