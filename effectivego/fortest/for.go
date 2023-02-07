package fortest

import (
	"fmt"
	"effectivego/iftest"
)

func SumFor(num int) int {
	sum := 0
	for i := 0; i <= num; i++ {
		sum += i
	}
	return sum
}
func SumWhile() int {
	sum := 0
	i := 0
	for i <= 10 {
		sum += i
		i++
	}
	return sum
}

func AbsIf(f []float64) []float64 {
	for k, v := range f {
		f[k] = iftest.AbsTest(v)
	}
	return f
}

func TestFor() {
	for pos, char := range "pinkman" {
		fmt.Printf("pos=%d,char=%d", pos, char)
	}
}
