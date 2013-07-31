package main

import (
	"html/template"
	"net/http"
	"net/url"
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
	routeNumber := len(routeNames)
	urls := make([]*url.URL, routeNumber)
	data.Links = make([]Link, routeNumber)
	for i := 0; i < routeNumber; i++ {
		urls[i], _ = rtr.Get(routeNames[i]).URL()
		data.Links[i] = Link{urls[i].String(), routeNames[i]}
	}
	t.Execute(w, data)
}
