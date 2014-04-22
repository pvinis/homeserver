package main

import (
	"html/template"
//	"github.com/gorilla/mux"
	"net/http"
//	"log"
)

type HomeTemplateData struct {
	BaseTemplateData
	Links map[string]string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.tmpl")
	data := HomeTemplateData{}
	data.Title = "home"

	data.Links = make(map[string]string)
	for _,n := range routes {
		data.Links[n] = router.Get(n).GetName()
	}

//	for k,v := range a {
//		log.Println(k, v)
//	}
	t.Execute(w, data)
}

