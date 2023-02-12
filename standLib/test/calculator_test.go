package test

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	if got != 5 {
		t.Errorf("2+3 should by 5 ,but get %d", got)
	}
}

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got < 0 {
		t.Errorf("Abs(-1) should be 1,but get %d", got)
		t.Name()
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(100)
	}
}

func BenchmarkAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Avg(100)
	}
}
