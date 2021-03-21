package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/kaakaa/mattermost-emojigen/server/font"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// EmojigenPlugin is the struct for mattermost-emojigen plugin
type EmojigenPlugin struct {
	plugin.MattermostPlugin
	client *MattermostClient
	UserID string

	configurationLock sync.RWMutex
	configuration     *configuration

	siteURL string
	router  *mux.Router

	drawer *font.EmojiDrawer
}

// OnActivate activate mattermost-emojigen plugin
func (p *EmojigenPlugin) OnActivate() error {
	p.API.LogInfo("Activating...")

	path, err := p.API.GetBundlePath()
	if err != nil {
		p.API.LogError("Failed to get bundle path")
		return err
	}
	drawer, err := font.NewEmojiDrawer(path)
	if err != nil {
		p.API.LogError("Failed to init EmojiDrawer", "details", err)
		return err
	}
	p.drawer = drawer

	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          "emojigen",
		AutoComplete:     true,
		AutoCompleteDesc: `Generate emoji`,
		AutoCompleteHint: `[EMOJI_NAME] [TEXT] [Black|Red|Blue|Green|White] [Black|Red|Blue|Green|White]`,
	}); err != nil {
		p.API.LogError(err.Error())
		return err
	}

	p.router = p.InitAPI()
	return nil
}

// ExecuteCommand handle commands that are created by this plugin
func (p *EmojigenPlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	if len(strings.Split(strings.TrimSpace(args.Command), " ")) == 1 {
		if appErr := p.API.OpenInteractiveDialog(p.getEmojiDialog(args.TriggerId)); appErr != nil {
			return nil, appErr
		}
		return &model.CommandResponse{}, nil
	}

	emoji, err := font.NewEmojiInfoFromLine(args.Command)
	if err != nil {
		appErr := &model.AppError{
			Message: fmt.Sprintf("Encountered error when parsing command: %v", err.Error()),
			StatusCode: http.StatusBadRequest,
			Where: "ExecuteCommand",
		}
		p.SendEphemeralPost(args.ChannelId, args.UserId, args.RootId, appErr.Error())
		return &model.CommandResponse{}, appErr
	}

	appErr := p.RegisterEmoji(emoji)
	if appErr != nil {
		p.SendEphemeralPost(args.ChannelId, args.UserId, args.RootId, appErr.Error())
		return &model.CommandResponse{}, appErr
	}
	p.SendEphemeralPost(args.ChannelId, args.UserId, args.RootId,fmt.Sprintf("Creating emoji with `:%s:` is success. :%s:", emoji.Name, emoji.Name))
	return &model.CommandResponse{}, nil
}

// OnConfigurationChange handle changes of configuration
func (p *EmojigenPlugin) OnConfigurationChange() error {
	var configuration = new(configuration)

	if err := p.API.LoadPluginConfiguration(configuration); err != nil {
		return fmt.Errorf("failed to load plugin configuration: %v", err)
	}

	p.setConfiguration(configuration)
	return p.setMattermostClient()
}

func (p *EmojigenPlugin) RegisterEmoji(emoji *font.EmojiInfo) *model.AppError {
	p.API.LogDebug(fmt.Sprintf("emoji: %#v", emoji))
	b, err := p.drawer.GenerateEmoji(emoji)
	if err != nil {
		return &model.AppError {
			Message: fmt.Sprintf("failed to generate emoji. details: %s", err.Error()),
			StatusCode: http.StatusBadRequest,
			Where: "RegisterEmoji",
		}
	}
	if err := p.client.RegistNewEmoji(b, emoji.Name, p.UserID); err != nil {
		return &model.AppError {
			Message: fmt.Sprintf("failed to register emoji. details: %s", err.Error()),
			StatusCode: http.StatusBadRequest,
			Where: "RegisterEmoji",
		}
	}
	return nil
}

func (p *EmojigenPlugin) setMattermostClient() error {
	if p.configuration == nil || p.configuration.AccessToken == "" {
		return fmt.Errorf("failed to load plugin configuration")
	}
	config := p.API.GetConfig()
	p.siteURL = *config.ServiceSettings.SiteURL

	c, err := Login(p.siteURL, p.configuration.AccessToken)
	if err != nil {
		return nil
	}
	id, err := c.getUserId()
	if err != nil {
		return err
	}

	p.client = c
	p.UserID = id

	return nil
}

func (p *EmojigenPlugin) getEmojiDialog(triggerID string) model.OpenDialogRequest {
	return model.OpenDialogRequest{
		TriggerId: triggerID,
		URL:       fmt.Sprintf("%s/plugins/%s/%s", p.siteURL, manifest.ID, "dialog/open"),
		Dialog: model.Dialog{
			Title: "Generate Emoji",
			Elements: []model.DialogElement{
				{
					DisplayName: "Emoji Name (e.g.: +1)",
					Name:        "emoji_name",
					Type:        "text",
					MinLength:   1,
					Placeholder: "+1, smiley,...",
				},
				{
					DisplayName: "Emoji Text",
					Name:        "emoji_text",
					Type:        "textarea",
					MinLength:   1,
					MaxLength:   20,
					HelpText:    "Text to be rendered as emoji. Text must be 1 ~ 20 characters including new line char.",
					Placeholder: "やば\nすぎ",
				},
				{
					DisplayName: "Font Color",
					Name:        "emoji_font_color",
					Type:        "select",
					Default:     "Black",
					Options: []*model.PostActionOptions{
						{Text: "Black", Value: "Black"},
						{Text: "Red", Value: "Red"},
						{Text: "Green", Value: "Green"},
						{Text: "Blue", Value: "Blue"},
						{Text: "White", Value: "White"},
					},
				},
				{
					DisplayName: "Background Color",
					Name:        "emoji_background_color",
					Type:        "select",
					Default:     "White",
					Options: []*model.PostActionOptions{
						{Text: "Black", Value: "Black"},
						{Text: "Red", Value: "Red"},
						{Text: "Green", Value: "Green"},
						{Text: "Blue", Value: "Blue"},
						{Text: "White", Value: "White"},
					},
				},
			},
			SubmitLabel: "Create",
		},
	}
}

// SendEphemeralPost sends an ephemeral post to a user as the bot account
func (p *EmojigenPlugin) SendEphemeralPost(channelID, userID, rootID, message string) {
	ephemeralPost := &model.Post{
		ChannelId: channelID,
		UserId:    p.UserID,
		RootId:    rootID,
		Message:   message,
	}
	_ = p.API.SendEphemeralPost(userID, ephemeralPost)
}
