package main

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

// MattermostClient is the client that sends requests to Mattermost REST APi
type MattermostClient struct {
	client *model.Client4
}

// Login setup MattermostClient
func Login(url, token string) (*MattermostClient, error) {
	c := model.NewAPIv4Client(url)
	c.AuthToken = token
	c.AuthType = model.HEADER_BEARER
	return &MattermostClient{
		client: c,
	}, nil
}

func (c *MattermostClient) getUserID() (string, error) {
	u, resp := c.client.GetMe("")
	if resp.Error != nil {
		return "", fmt.Errorf("failed to get user id: %w", resp.Error)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get user id: %d", resp.StatusCode)
	}
	return u.Id, nil
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
