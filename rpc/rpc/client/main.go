package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	p := Params{5, 6}
	ret := 0
	err = conn.Call("Rect.Area", p, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("面积：", ret)
	// 周长
	err = conn.Call("Rect.Perimeter", p, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长：", ret)
}
