package internal

// QuickSort 快速排序，根据 field 字段对map数组倒序排序
func QuickSort(m []map[string]string, left, right int, field string) {
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
		for i < j && m[i][field] > tmp[field] {
			i++
		}
		m[j] = m[i]
	}
	m[i] = tmp
	QuickSort(m, left, i-1, field)
	QuickSort(m, i+1, right, field)
}
