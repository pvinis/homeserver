package main

import (
        "html/template"
        "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("templates/home.tmpl")
        data := BaseTemplate{"home"}
        t.Execute(w, data)
}