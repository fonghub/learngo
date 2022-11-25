package structs

// Construct 构造器
func Construct(name string, age int, nickname string) Persons {
	return Persons{name, age, nickname}
}
