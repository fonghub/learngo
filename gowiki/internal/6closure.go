package internal

import (
	"fmt"
	"net/http"
	"os"
)

// MakeHandler 闭包是一个使用 MakeHandler 返回的，携带http.ResponseWriter和http.Request参数的函数
// 闭包包含有定义于外部的变量
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	var p1 = new(Page)
	bookInfo, err := p1.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	content, _ := os.ReadFile("./html/book.html")
	fmt.Fprintf(w, string(content), title, bookInfo["name"], bookInfo["author"], bookInfo["price"])
}
