package internal

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type form struct {
}

// Form 使用html模版实现表单的显示，修改，保存
var Form = new(form)

func (f *form) ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	var p1 = new(Page)
	p, err := p1.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	p["title"] = title
	f.RenderTemplate(w, "view", p)
}

func (f *form) RenderTemplate(w http.ResponseWriter, tmpl string, p map[string]interface{}) {
	t, err := template.ParseFiles("html/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (f *form) EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	var p1 = new(Page)
	p, err := p1.LoadPage(title)
	if err != nil {
		p = map[string]interface{}{"Title": title, "author": "", "name": "", "price": ""}
	}
	p["title"] = title
	f.RenderTemplate(w, "edit", p)
}

func (f *form) SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	var data = make(map[string]interface{})
	data["author"] = r.FormValue("author")
	data["name"] = r.FormValue("name")
	data["price"] = r.FormValue("price")
	jsn, _ := json.Marshal(data)
	p := &Page{Title: title, Body: []byte(string(jsn))}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
