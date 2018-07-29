package main

import (
	"log"

	"github.com/kaakaa/mattermost-emojigen/util"
)

func main() {
	url := "http://example.com"
	token := "3pxeuyjd8t8a9fmny4myeby1da"
	client := util.Login(url, token)
	log.Println("Get Mattermost client")
	if err := client.RegistNewEmoji("hoge", "hoge", "user_id"); err != nil {
		log.Fatalf("Regist Error: %v", err)
	}
}
