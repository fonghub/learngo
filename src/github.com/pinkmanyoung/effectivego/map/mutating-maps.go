package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 43
	fmt.Println("the value:",m["Answer"])

	m["Answer"] = 48
	fmt.Println("the value:",m["Answer"])

	v,ok := m["Answer"]
	fmt.Println("the value:",v,"Present?",ok)
	
	delete(m,"Answer")
	fmt.Println("the value:",m["Answer"])
}