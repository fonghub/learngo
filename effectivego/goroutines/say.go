package goroutines

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, "=", i)
	}
}

// Seq 协程的调度不保证执行的顺序
func Seq(s string) {
	for i := 0; i < 5; i++ {
		go fmt.Println(s, "=", i)
	}
}
