package structs

import "fmt"

type first struct{}

func (f *first) gd() {
	fmt.Println("去 gd 旅游")
}
func (f *first) gx() {
	fmt.Println("去 gx 旅游")
}

type second struct {
	first
}

func (s *second) yn() {
	fmt.Println("去 yn 旅游")
}
func (s *second) sc() {
	fmt.Println("去 sc 旅游")
}

type third struct {
	second
}

func (t *third) hn() {
	fmt.Println("去 hn 旅游")
}
func (t *third) hb() {
	fmt.Println("去 hb 旅游")
}

type fouth struct {
	third
}

func (f *fouth) fj() {
	fmt.Println("去 fj 旅游")
}
func (f *fouth) xj() {
	fmt.Println("去 xj 旅游")
}

// FourYear 结构体的隐式继承
func FourYear() {
	fourYear := new(fouth)
	fourYear.first.gd()
	fourYear.gx()

	fourYear.second.yn()
	fourYear.sc()

	fourYear.third.hn()
	fourYear.hb()

	fourYear.fj()
	fourYear.xj()
}
