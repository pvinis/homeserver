package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os/exec"
)

var rtr = mux.NewRouter()

func main() {
	rtr.HandleFunc("/", HomeHandler).Name("home")
	rtr.HandleFunc("/test", TestHandler).Name("test")
	http.Handle("/", rtr)
	http.ListenAndServe(":8001", nil)
}

type Page struct {
	Title string
	Name  string
	Url   string
}

type LinksPage struct {
	Page
	Links []Link
}

type TestPage struct {
	Page
	Names []string
}

type Link struct {
	Url  string
	Name string
}

func loadPage(name string) {

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	data := LinksPage{}
	data.Title = " aa"
	u1, _ := rtr.Get("test").URL()
	data.Links = []Link{Link{u1.String(), "test"}, Link{"2", "d"}, Link{"3", "asd"}}
	t.Execute(w, data)
}

/// pass
func TestHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("test.html")
	data := TestPage{}
	data.Title = "test"
	u1, _ := rtr.Get("test").URL()
	data.Url = u1.String()
	names, err := exec.Command("ls").Output()
	if err != nil {
		println("shit")
	}
	data.Names = []string{string(names), "Asd"}
	t.Execute(w, data)
	////regfresh
}
