package main

import (
	"github.com/gorilla/mux"
	///        "os/exec"
)

var router = mux.NewRouter()
var routes = []string{"test"}

func RegisterHandlers() {
	router.Path("/").HandlerFunc(HomeHandler).Name("home")
	//router.HandleFunc("/notify", NotifyHandler).Name("notify")
	router.Path("/test").HandlerFunc(TestHandler).Name("test")
	///      rtr.HandleFunc("/torrents", TorrentsHandler).Name("torrents")
	///   rtr.HandleFunc("/status", StatusHandler).Name("status")
}
