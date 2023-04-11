package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
)

type vals struct {
	x int
	y int
}

// MakeHomeWork 加减法生成工具
func MakeHomeWork(num int, max int) {
	var sm sync.Map
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			x := rand.Intn(max)
			y := rand.Intn(max)
			res := x + y
			sm.Store(vals{x, y}, res)
			wg.Done()
		}()
	}
	wg.Wait()
	var ii = 1
	sm.Range(func(key, value any) bool {
		if ii%4 != 0 {
			fmt.Printf("%v+%v=\t\t", key.(vals).x, key.(vals).y)
		} else {
			fmt.Printf("%v+%v=\n", key.(vals).x, key.(vals).y)
		}
		ii++
		return true
	})
	fmt.Println("-------------")
	ii = 1
	sm.Range(func(key, value any) bool {
		if ii%4 != 0 {
			fmt.Printf("%v-%v=\t\t", key.(vals).x+key.(vals).y, key.(vals).y)
		} else {
			fmt.Printf("%v-%v=\n", key.(vals).x+key.(vals).y, key.(vals).y)
		}
		ii++
		return true
	})
}
