package test

func Add(a, b int) int {
	return a + b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func Sum(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}

func Avg(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			total += i
		}
	}
	return total
}
