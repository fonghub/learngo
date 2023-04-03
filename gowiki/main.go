package main

import (
	"encoding/json"
	"fmt"
	"gowiki/internal"
	"log"
	"net/http"
)

func main() {
	//testReadingFile()
	//testNetHttp()
	//testNetHttpPage()
	//testNetHttpForm()
	//testValidation()
	testClosure()
}

func testReadingFile() {
	var data = make(map[string]interface{})
	data["name"] = "TGPL"
	data["author"] = "james"
	data["price"] = 22.3
	jsn, _ := json.Marshal(data)
	p1 := &internal.Page{Title: "TGPL", Body: []byte(string(jsn))}
	p1.Save()
	data2, _ := p1.LoadPage(p1.Title)
	fmt.Println("data2", data2)
}

// testNetHttp http://localhost:8080/golang
func testNetHttp() {
	http.HandleFunc("/", internal.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// testNetHttpPage http://localhost:8080/view/tgpl
func testNetHttpPage() {
	http.HandleFunc("/view/", internal.Pg.ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testNetHttpForm() {
	http.HandleFunc("/view/", internal.Form.ViewHandler)
	http.HandleFunc("/edit/", internal.Form.EditHandler)
	http.HandleFunc("/save/", internal.Form.SaveHandler)
	http.ListenAndServe(":8080", nil)
}

func testValidation() {
	http.HandleFunc("/view/", internal.Get)
	http.ListenAndServe(":8080", nil)
}

func testClosure() {
	http.HandleFunc("/404/", internal.Handler)
	http.HandleFunc("/view/", internal.MakeHandler(internal.ViewHandler))
	http.ListenAndServe(":8080", nil)
}
