package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaakaa/mattermost-emojigen/server/font"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

var infoMessage = "This is Mattermost Emojigen v" + manifest.Version + "\n"

// InitAPI returns a router for mattermost-emojigen plugin
func (p *EmojigenPlugin) InitAPI() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", p.handleInfo).Methods("GET")

	dialogRouter := r.PathPrefix("/dialog").Subrouter()
	dialogRouter.HandleFunc("/open", p.handleSubmitDialog).Methods("POST")
	return r
}

func (p *EmojigenPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("New request:", "Host", r.Host, "RequestURI", r.RequestURI, "Mehotd", r.Method)
	p.router.ServeHTTP(w, r)
}

func (p *EmojigenPlugin) handleInfo(w http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(w, infoMessage)
}

func (p *EmojigenPlugin) handleSubmitDialog(w http.ResponseWriter, r *http.Request) {
	request := model.SubmitDialogRequestFromJson(r.Body)
	if request == nil {
		p.API.LogWarn("Failed to decode SubmitDialogRequest")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	emojiInfo, err := requestToEmojiInfo(request)
	if err != nil {
		p.API.LogWarn("Failed to parse SubmitDialogRequest", "details", err.Error())
		http.Error(w, "invalid request", http.StatusBadRequest)
		p.SendEphemeralPost(request.UserId, request.ChannelId, "", fmt.Sprintf("failed to create emoji. details: %s", err.Error()))
		return
	}
	appErr := p.RegisterEmoji(emojiInfo)
	if appErr != nil {
		p.API.LogWarn("Failed to create emoji.", "details", appErr.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		p.SendEphemeralPost(request.UserId, request.ChannelId, "", fmt.Sprintf("failed to create emoji. details: %s", appErr.Error()))
		return
	}

	p.SendEphemeralPost(request.UserId, request.ChannelId, "", fmt.Sprintf("creating emoji with `:%s:` is success. :%s:", emojiInfo.Name, emojiInfo.Name))
	w.WriteHeader(http.StatusOK)
}

func requestToEmojiInfo(request *model.SubmitDialogRequest) (*font.EmojiInfo, error) {
	name, ok := request.Submission["emoji_name"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get `emoji_name` from submission. %#v", request.Submission)
	}
	text, ok := request.Submission["emoji_text"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get `emoji_text` from submission. %#v", request.Submission)
	}
	fc, ok := request.Submission["emoji_font_color"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get `emoji_font_color` from submission. %#v", request.Submission)
	}
	fontColor, err := font.ColorFromString(fc)
	if err != nil {
		return nil, err
	}

	bgc, ok := request.Submission["emoji_background_color"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get `emoji_background_color` from submission. %#v", request.Submission)
	}
	bgColor, err := font.ColorFromString(bgc)
	if err != nil {
		return nil, err
	}

	return &font.EmojiInfo{
		Name:            name,
		Text:            text,
		FontColor:       fontColor,
		BackgroundColor: bgColor,
	}, nil
}
