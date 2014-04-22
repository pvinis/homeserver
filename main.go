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
		err = http.ListenAndServe(":3001", nil)
		if err != nil {
			 //////logger
			 fmt.Println("something is wrong")
		}
	} else {
		fmt.Println("Serving on port 80..")
	}
}
