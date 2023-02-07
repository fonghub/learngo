package iftest

import "math"

func AbsTest(x float64) float64 {
	if x < 0 {
		return math.Abs(x)
	}
	return x
}

func PowTest(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
