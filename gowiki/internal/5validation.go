package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func Get(w http.ResponseWriter, r *http.Request) {
	s, e := GetTitle(w, r)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("s", s)
}
func GetTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	fmt.Println("m", m)
	return m[2], nil
}
