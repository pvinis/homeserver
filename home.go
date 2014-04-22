package main

import (
	"html/template"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

type HomeTemplateData struct {
	BaseTemplateData
	links []string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.tmpl")
	data := HomeTemplateData{}
	data.Title = "home"
	var a map[string]*mux.Route
	//a = router.getNamedRoutes()
	a["asd"] = router.NewRoute()
	for k,v := range a {
		fmt.Println(k, v)
	}
	t.Execute(w, data)
}

