package main

import (
	"html/template"
	"net/http"
	"bitbucket.org/kisom/gopush/pushover"
	"os/exec"
	"regexp"
)

type HomeTemplateData struct {
	BaseTemplateData
	ServerAlive bool
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.tmpl")

	data := HomeTemplateData{}
	data.Title = "home"
	data.ServerAlive = pingServer()
	t.Execute(w, data)
}

func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	sendNotification()
	http.Redirect(w, r, "/?notified", http.StatusFound)
}

func pingServer() bool {
	re := regexp.MustCompile("(\\d+) received")
	out, err := exec.Command("ping", "-c", "1", "192.168.2.153").Output()
	if err != nil {
		/////log it	fmt.Println(err)
	}
	packets := re.ReplaceAllString(re.FindString(string(out)), "$1")
	if packets == "1" {
		return true
	}
	return false
}

func sendNotification() bool {
	identity := pushover.Authenticate("ajUCEgKvd9QmeMKFHBmvwMct1DYytx","iBzANa9RG4jhS3s9QK9MAGNruSoGSz")
	sent := pushover.Notify(identity, "turn on server")
	if !sent {
	///log
	return false
	}
return true
}

