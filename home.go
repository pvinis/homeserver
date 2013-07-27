package main

import (
	"html/template"
	"net/http"
)

type LinksPage struct {
	Page
	Links []Link
}

type Link struct {
	Url  string
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	data := LinksPage{}
	data.Title = "Home"
	u1, _ := rtr.Get("test").URL()
	u2, _ := rtr.Get("torrents").URL()
	data.Links = []Link{Link{u1.String(), "test"}, Link{u2.String(), "torrents"}, Link{"3", "transmission"}}
	t.Execute(w, data)
}
