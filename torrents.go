package main

import (
	"html/template"
	"net/http"
	"os/exec"
)

type TorrentStatus int

const (
	Running TorrentStatus = iota
	Paused
	Finished
)

type Torrent struct {
	Name     string
	Status   TorrentStatus
	DLSpeed  int
	ULSpeed  int
	Progress int
}

type TorrentPage struct {
	Page
	Torrents []Torrent
}

func TorrentsHandler(w http.ResponseWriter, r *http.Request) {
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
