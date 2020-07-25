package main

import (
	"encoding/json"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var config = []Config{}

func parseGhPost(rw http.ResponseWriter, request *http.Request) {
	path := strings.TrimLeft(request.URL.Path, "/")
	if request.Method != "POST" {
		return
	}
	if request.Body == http.NoBody {
		return
	}

	var message string = ""
	var webhookUrl string = ""
	var channel string = ""
	for _, v := range config {
		if v.Name == path {
			webhookUrl = v.Options.Webhook
			channel = v.Options.Channel
			fmt.Println("Settings found")
		}
	}
	if len(strings.TrimSpace(webhookUrl)) == 0 {
		fmt.Println("Exit")
	}
	decoder := json.NewDecoder(request.Body)
	var t WebHook
	err := decoder.Decode(&t)

	if err != nil {
		return
	}

	if len(t.Event.Exception.Values) > 0 {
		message = t.Event.Exception.Values[0].Value
	} else {
		message = t.Message
	}

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Project", Value: t.Event.Environment + " " + t.Culprit})
	attachment1.AddField(slack.Field{Title: "Error", Value: message})
	payload := slack.Payload{
		Text:        "Sentry error " + t.URL,
		Username:    "sentry",
		Channel:     channel,
		IconEmoji:   ":aws:",
		Attachments: []slack.Attachment{attachment1},
	}
	slack.Send(webhookUrl, "", payload)

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {

	filename, _ := filepath.Abs("./config.yaml")

	if fileExists(filename) {
		fmt.Println("Config file exists")
	} else {
		fmt.Println("Config file does not exist (or is a directory)")
		os.Exit(0)
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(0)
	}
	err = yaml.Unmarshal(yamlFile, &config)

	for i, s := range config {
		fmt.Println(i, s.Options)
		http.HandleFunc("/"+s.Options.Path, parseGhPost)

	}
	http.ListenAndServe(":8080", nil)

}
