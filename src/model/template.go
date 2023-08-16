package model

import (
	"html/template"
	"net/http"
)

type Page struct {
	Path string
	Data any
}

func (p *Page) ShowTemplate(w http.ResponseWriter) {
	tml := template.Must(template.ParseFiles("template/" + p.Path))
	tml.Execute(w, p.Data)
}
