package internal

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"os"
)

type page struct {
}

var Pg = new(page)

// ViewHandler 从url截取参数，从文件加载指定文件，获取参数，渲染到模板
func (pg *page) ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	var p1 = new(Page)
	p, _ := p1.LoadPage(title)
	content, _ := os.ReadFile("./html/book.html")
	fmt.Fprintf(w, string(content), title, gconv.String(p["name"]), gconv.String(p["author"]), gconv.String(p["price"]))
}
