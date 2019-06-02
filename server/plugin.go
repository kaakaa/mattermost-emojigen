package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/kaakaa/mattermost-emojigen/server/font"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

type EmojigenPlugin struct {
	plugin.MattermostPlugin
	client *MattermostClient

	configurationLock sync.RWMutex
	configuration     *configuration

	drawer *font.EmojiDrawer
}

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

	b, err := p.drawer.GenerateEmoji(emojiText)

	p.API.LogDebug(fmt.Sprintf("TEST: %v", len(b)))
	if err != nil {
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Encountered error when generating emoji image: %v", err.Error()),
		}, nil
	}

	if err := p.client.RegistNewEmoji(b, emojiName, userId); err != nil {
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
	var configuration = new(configuration)

	if err := p.API.LoadPluginConfiguration(configuration); err != nil {
		return fmt.Errorf("failed to load plugin configuration: %v", err)
	}

	p.setConfiguration(configuration)
	return p.setMattermostClient()
}

func (p *EmojigenPlugin) setMattermostClient() error {
	if p.configuration == nil || p.configuration.AccessToken == "" {
		return fmt.Errorf("failed to load plugin configuration")
	}
	config := p.API.GetConfig()
	p.client = Login(*config.ServiceSettings.SiteURL, p.configuration.AccessToken)
	p.API.LogInfo(fmt.Sprintf("SiteURL: %v", *config.ServiceSettings.SiteURL))
	p.API.LogInfo(fmt.Sprintf("AccessToken: %v", p.configuration.AccessToken))

	return nil
}
