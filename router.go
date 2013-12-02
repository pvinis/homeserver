package main

import (
	"github.com/gorilla/mux"
	///        "os/exec"
)

var router = mux.NewRouter()

func RegisterHandlers() {
	router.HandleFunc("/", HomeHandler).Name("home")
	router.HandleFunc("/notify", NotifyHandler).Name("notify")
	///        rtr.HandleFunc("/test", TestHandler).Name("test")
	///      rtr.HandleFunc("/torrents", TorrentsHandler).Name("torrents")
	///   rtr.HandleFunc("/status", StatusHandler).Name("status")
}
