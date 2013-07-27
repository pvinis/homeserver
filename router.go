package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os/exec"
)

var rtr = mux.NewRouter()

func Init() {
	rtr.HandleFunc("/", HomeHandler).Name("home")
	rtr.HandleFunc("/test", TestHandler).Name("test")
	rtr.HandleFunc("/torrents", TorrentsHandler).Name("torrents")
}

type Page struct {
	Title string
	Name  string
	Url   string
}

type TestPage struct {
	Page
	Names []string
}

func loadPage(name string) {
	//////
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
