package main

import (
	"fmt"
	"strings"

	"github.com/kaakaa/mattermost-emojigen/util"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

type EmojigenPlugin struct {
	plugin.MattermostPlugin
	client      *util.MattermostClient
	AccessToken string
}

func (p *EmojigenPlugin) OnActivate() error {
	p.API.LogInfo("Activating...")

	p.setMattermostClient()

	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          "emojigen",
		AutoComplete:     true,
		AutoCompleteDesc: `Generate emoji`,
		AutoCompleteHint: `[EMOJI_NAME] [TEXT]`,
	}); err != nil {
		p.API.LogError(err.Error())
		return err
	}
	return nil
}

func (p *EmojigenPlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	text := strings.Split(args.Command, " ")
	emojiName := text[1]
	emojiText := text[2]
	userId := args.UserId
	p.API.LogDebug(fmt.Sprintf("emoji_name: %v", emojiName))
	p.API.LogDebug(fmt.Sprintf("message: %v", emojiText))
	p.API.LogDebug(fmt.Sprintf("user_id: %v", userId))

	if err := p.client.RegistNewEmoji(emojiName, emojiText, userId); err != nil {
		p.API.LogError(err.Error())
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Encountered error when creating emoji: %v", err.Error()),
		}, nil
	}
	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         fmt.Sprintf("Creating emoji with `:%s:` is success. :%s:", emojiName, emojiName),
	}, nil
}

func (p *EmojigenPlugin) OnConfigurationChange() error {
	if err := p.MattermostPlugin.OnConfigurationChange(); err != nil {
		p.API.LogError(err.Error())
		return err
	}

	p.setMattermostClient()
	return nil
}

func (p *EmojigenPlugin) setMattermostClient() {
	config := p.API.GetConfig()
	p.client = util.Login(*config.ServiceSettings.SiteURL, p.AccessToken)
	p.API.LogInfo(fmt.Sprintf("Update client successfuly"))
	p.API.LogInfo(fmt.Sprintf("SiteURL: %v", *config.ServiceSettings.SiteURL))
	p.API.LogInfo(fmt.Sprintf("AccessToken: %v", p.AccessToken))
}

func main() {
	plugin.ClientMain(&EmojigenPlugin{})
}
