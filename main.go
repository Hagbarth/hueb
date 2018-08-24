package main

import (
	"log"
	"net/http"

	"github.com/Hagbarth/hueb/api"
	"github.com/Hagbarth/hueb/ifttt"
	"github.com/Hagbarth/hueb/lights"
	"github.com/Hagbarth/hueb/webapp"
	"github.com/namsral/flag"
)

func main() {
	flag.String(flag.DefaultConfigFlagname, "go.conf", "path to config file")
	iftttKey := flag.String("ifttt-key", "", "your ifttt api key")
	flag.Parse()

	http.Handle("/", webapp.NewServer())

	iftttClient := ifttt.NewClient(*iftttKey)
	switcher := lights.NewSwitcher(iftttClient)
	http.Handle("/api/", http.StripPrefix("/api", api.NewServer(switcher)))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
