package main

import (
        "net/http"
		"fmt"
)

func main() {
        RegisterHandlers()
        http.Handle("/", router)
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			err = http.ListenAndServe(":3000", nil)
			if err != nil {
				fmt.Println(err)//////logger
			}
		} else {
			fmt.Println("Serving on port 80..")
		}
}
