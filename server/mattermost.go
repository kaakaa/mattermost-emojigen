package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-server/model"
)

type MattermostClient struct {
	client *model.Client4
}

func Login(url, token string) *MattermostClient {
	c := model.NewAPIv4Client(url)
	c.AuthToken = token
	c.AuthType = model.HEADER_BEARER
	return &MattermostClient{
		client: c,
	}
}

// func (c *MattermostClient) RegistNewEmoji(name, msg, userId string) error {
func (c *MattermostClient) RegistNewEmoji(b []byte, name, userId string) error {
	_, resp := c.client.CreateEmoji(&model.Emoji{
		CreatorId: userId,
		Name:      name,
	}, b, "emojigen")

	log.Printf("Response from Mattermost: %#v", resp)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Error.Message)
	}
	return nil
}

func writeResponse(w http.ResponseWriter, response model.CommandResponse) {
	if _, err := io.WriteString(w, response.ToJson()); err != nil {
		log.Printf(fmt.Sprintf("Error: Cannot response successfuly. err=%v", err.Error()))
		return
	}
}
