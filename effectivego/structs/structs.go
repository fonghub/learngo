package structs

type vertex struct {
	X, Y int
}

type Persons struct {
	Name     string
	Age      int
	Nickname string
}

// Value 普通方式访问结构体字段
func Value() vertex {
	v := vertex{2, 2}
	return v
}

// Point 指针方式访问结构体字段
func Point() *vertex {
	v := &vertex{3, 4}
	return v
}
