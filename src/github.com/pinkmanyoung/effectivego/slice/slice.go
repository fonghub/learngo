package main

import "fmt"

func main(){
	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
	var s []int = primes[1:4]
	test(s)
	fmt.Println(s)
	fmt.Println(primes)
}

func test(s []int){
	fmt.Println(s)
	for k,v := range s {
		fmt.Printf("k=%d,v=%d\n",k,v)
	}
	s[0] = 11;
}