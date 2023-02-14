package maps

import (
	"fmt"
	"strings"
	"time"
)

type vertex struct {
	lat, long float64
}

// MapTest 创建map
func MapTest() map[string]vertex {
	//创建map结构
	m := make(map[string]vertex)
	//对map赋值
	m["Bell Labs"] = vertex{
		40.4, -74.39,
	}
	return m
}

// MapValue 创建map
func MapValue() map[string]vertex {
	//创建map，同时对map赋值
	m := map[string]vertex{
		"Bell Labs": {40.4, -74.39}, //省略了vertex类型关键字
		"Google":    {37.42202, -122.08408},
	}
	return m
}

// MapRange 遍历map
func MapRange(m map[string]vertex) {
	for k, v := range m {
		fmt.Printf("%v's vertex:[lat:%v,long:%v]%v\n", k, v.lat, v.long)
	}
}

func MapCURD() {
	//创建一个空map
	m := map[string]vertex{}
	fmt.Println(m)
	//添加/修改map元素
	m["Bell Labs"] = vertex{40.4, -74.39}
	//获取map元素
	fmt.Println(m["Bell Labs"])
	//删除map元素
	delete(m, "Bell Labs")
	if elem, ok := m["Bell Labs"]; ok {
		//获取map元素
		fmt.Println(elem)
	} else {
		fmt.Println("该键不存在")
	}

	//打印map
	fmt.Println(m)
}

func MapExercise(s string) map[string]int {
	arr := strings.Fields(s)
	m := map[string]int{}
	for _, v := range arr {
		m[v]++
	}
	return m
}

// MapSearch 遍历map，输出顺序不确定
func MapSearch() {
	var m = map[string]string{
		"game_id":     "17",
		"create_by":   "22",
		"game_name":   "slot",
		"status":      "1",
		"create_time": time.Now().String(),
	}
	fmt.Println("m", m)
	for k, v := range m {
		fmt.Printf("k=%s,v=%s\n", k, v)
	}
}
