package main

import (
	"net/http"
)

func main() {
	Init()
	http.Handle("/", rtr)
	http.ListenAndServe(":8001", nil)
}
