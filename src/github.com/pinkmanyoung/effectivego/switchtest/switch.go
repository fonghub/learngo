package switchtest

import (
	"fmt"
	"runtime"
)

func OsTest() string {
	fmt.Print("Go runs on ")
	var osname string
	switch os := runtime.GOOS; os {
	case "windows":
		osname = "windows os"
	case "darwin":
		osname = "OS X"
	case "linux":
		osname = "linux os"
	default:
		osname = "other os"
	}
	return osname
}

func AdjAge(age int) string {
	var str string
	//switch关键字后没有跟判断值，则case后面要跟表达式
	switch {
	case age < 10:
		str = "少年"
	case age < 18:
		str = "未成年人"
	case age < 50:
		str = "中年"
	case age < 30:
		str = "青年"
	}
	return str
}
