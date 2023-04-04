package main

import (
	"fmt"
	"sorting/internal"
)

func main() {
	testQuickSort()
}

func getData() []map[string]string {
	return []map[string]string{
		{"createTime": "2023-1-2", "gameId": "1002"},
		{"createTime": "2023-1-12", "gameId": "1012"},
		{"createTime": "2023-1-3", "gameId": "1003"},
		{"createTime": "2023-1-1", "gameId": "1001"},
		{"createTime": "2023-1-11", "gameId": "1011"},
		{"createTime": "2023-1-4", "gameId": "1004"},
	}
}
func testQuickSort() {
	list := getData()
	fmt.Println(list)
	internal.QuickSort(list, 0, len(list)-1, "gameId")
	fmt.Println("after sorting")
	fmt.Println(list)
}
