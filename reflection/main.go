package main

import (
	"reflection/internal"
)

func main() {
	testBegin()
	//testThreeLaws()
}

func testBegin() {
	internal.Demo1()
}
func testThreeLaws() {
	//internal.FirstLaw()
	//internal.SecondLaw()
	//internal.ThirdLaw()
	internal.ReadAndWriteStruct()
}
