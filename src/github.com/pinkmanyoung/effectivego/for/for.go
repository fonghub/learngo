package main

import "fmt"

func main(){
	for pos,char := range "pinkman" {
		fmt.Printf("pos=%d,char=%d",pos,char)
	}
}