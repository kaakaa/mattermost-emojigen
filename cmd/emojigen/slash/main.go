package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kaakaa/mattermost-emojigen/util"
)

var config = flag.String("c", "config.json", "optional path to the config file")

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c, err := util.LoadConf(*config)
	if err != nil {
		log.Fatal(err)
	}
	mmClient := util.Login(c.MattermostHost, c.Token)
	http.HandleFunc("/emoji", mmClient.Server)
	if err := http.ListenAndServe(c.Listen, nil); err != nil {
		log.Fatal(err)
	}
}
