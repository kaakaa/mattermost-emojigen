package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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

func (c *MattermostClient) RegistNewEmoji(name, msg, userId string) error {
	_, resp := c.client.GetUser(userId, "")
	if len(userId) == 0 || resp.StatusCode != 200 {
		u, resp := c.client.GetMe("")
		if resp.StatusCode != 200 {
			return fmt.Errorf(resp.Error.Message)
		}
		userId = u.Id
	}

	b, err := generate(msg)
	if err != nil {
		return err
	}

	log.Printf("Generate Emoji: %s: %d bytes", msg, len(b))
	_, resp = c.client.CreateEmoji(&model.Emoji{
		CreatorId: userId,
		Name:      name,
	}, b, "emojigen")

	log.Printf("Response from Mattermost: %d", resp.StatusCode)
	if resp.StatusCode != 200 {
		return fmt.Errorf(resp.Error.Message)
	}
	return nil
}

func (c *MattermostClient) Server(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	var response model.CommandResponse
	response.Username = "Mattermost-Emojigen"
	response.ResponseType = model.COMMAND_RESPONSE_TYPE_EPHEMERAL

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		response.Text = fmt.Sprintf("System Error: Invalid Content-Type (%v) is received. Please contact administrator.", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		writeResponse(w, response)
		return
	}
	if err := r.ParseForm(); err != nil {
		response.Text = fmt.Sprintf("System Error: Cannot parse form parameters. err=%v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, response)
		return
	}

	m := r.Form
	text := strings.Split(m["text"][0], " ")
	emojiName := text[0]
	emojiText := text[1]
	userId := m["user_id"]
	if err := c.RegistNewEmoji(emojiName, emojiText, userId[0]); err != nil {
		response.Text = fmt.Sprintf("Error: Cannot create emoji. err=%v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, response)
		return
	}

	response.Text = fmt.Sprintf("Creating emoji with `:%s:` is success. :%s:", emojiName, emojiName)
	w.WriteHeader(http.StatusOK)
	writeResponse(w, response)
}

func writeResponse(w http.ResponseWriter, response model.CommandResponse) {
	if _, err := io.WriteString(w, response.ToJson()); err != nil {
		log.Printf(fmt.Sprintf("Error: Cannot response successfuly. err=%v", err.Error()))
		return
	}
}
