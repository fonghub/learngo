package main

import "fmt"

type Vertex struct {
	Lat,Long float64
}

var m = map[string]Vertex {
	"Bell Labs":Vertex{
		40.22,-70.99,
	},
	"Google":Vertex {
		37.33,-122.22,
	},
}

var mi = map[int]string {
	0:"Tom",
	1:"peter",
	2:"james",
}

func main() {
	fmt.Println(m)
	fmt.Println(mi)
}