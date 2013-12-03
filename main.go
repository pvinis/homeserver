package main

import (
	"fmt"
	"net/http"
)

func main() {
	RegisterHandlers()
	http.Handle("/", router)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		err = http.ListenAndServe(":3000", nil)
		if err != nil {
			 //////logger
		}
	} else {
		fmt.Println("Serving on port 80..")
	}
}
