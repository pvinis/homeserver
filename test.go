package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func msdfain() {
	names, err := exec.Command("transmission-remote", "-l").Output()

	if err != nil {
		println("shit")
	}
	n := string(names)
	fmt.Println(n)
	var a int
	var t []byte
	a, t, err = bufio.ScanLines(names[1:], false)
	fmt.Println(a, t, err)
}
