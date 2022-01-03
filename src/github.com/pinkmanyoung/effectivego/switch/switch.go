package main

import "fmt"

func adjAge(age int) string{
	switch {
		case age < 10:
			return "少年"
		case age < 18:
			return "未成年人"
		case age < 30:
			return "青年"
		case age < 50:
			return "中年"
	}
	return "老年人"
}

func main(){
	age := 102
	res := adjAge(age);
	fmt.Printf("age=%d,res=%s",age,res)
}

