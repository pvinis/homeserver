package main

import (
	"html/template"
	"net/http"
)

type StatusPage struct {
	Page
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("test.html")
	data := TestPage{}
	data.Title = "test"
	u1, _ := rtr.Get("test").URL()
	data.Url = u1.String()
	t.Execute(w, data)
	////regfresh
}
