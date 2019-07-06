package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mattermost/mattermost-server/model"
)

// MattermostClient is the client that sends requests to Mattermost REST APi
type MattermostClient struct {
	client *model.Client4
}

// Login setup MattermostClient
func Login(url, token string) *MattermostClient {
	c := model.NewAPIv4Client(url)
	c.AuthToken = token
	c.AuthType = model.HEADER_BEARER
	return &MattermostClient{
		client: c,
	}
}

// RegistNewEmoji send a request for creating emoji
func (c *MattermostClient) RegistNewEmoji(b []byte, name, userID string) error {
	_, resp := c.client.CreateEmoji(&model.Emoji{
		CreatorId: userID,
		Name:      name,
	}, b, "emojigen")

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Error.Message)
	}
	return nil
}

func writeResponse(w http.ResponseWriter, response model.CommandResponse) {
	if _, err := io.WriteString(w, response.ToJson()); err != nil {
		return
	}
}
