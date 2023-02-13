package internal

import "fmt"

func SliceMapSort() {
	list := []map[string]string{
		{"createTime": "2023-1-2", "gameId": "1002"},
		{"createTime": "2023-1-12", "gameId": "1012"},
		{"createTime": "2023-1-3", "gameId": "1003"},
		{"createTime": "2023-1-1", "gameId": "1001"},
		{"createTime": "2023-1-11", "gameId": "1011"},
		{"createTime": "2023-1-4", "gameId": "1004"},
	}
	fmt.Println(list)
	SliceMap(list, 0, len(list)-1, "gameId")
	fmt.Println(list)
}

// SliceMap 快速排序map数组 倒序
func SliceMap(m []map[string]string, left, right int, field string) {
	i := left
	j := right

	if left >= right {
		return
	}

	tmp := m[i]
	for i < j {
		for i < j && m[j][field] <= tmp[field] {
			j--
		}
		m[i] = m[j]
		for i < j && m[i][field] >= tmp[field] {
			i++
		}
		m[j] = m[i]
	}
	m[i] = tmp
	SliceMap(m, left, i-1, field)
	SliceMap(m, i+1, right, field)
}
