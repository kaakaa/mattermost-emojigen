[![Build Status](https://travis-ci.com/kaakaa/mattermost-emojigen.svg?token=Quh1sKztRpq9F7nwVCUC&branch=master)](https://travis-ci.com/kaakaa/mattermost-emojigen)

# Mattermost EmojiGen

Mattermost plugin for generating custom emoji.

![sample](./emoji_sample.png)

## Set up

### Plugin
1. Download a plugin distribution from [Releases · kaakaa/mattermost\-cron\-plugin](https://github.com/kaakaa/mattermost-emojigen/releases)
2. Upload and Enabling plugin from your mattermost's admin console
3. Set your Mattermost Personal Access Token in configuration page on admin console

**Warning**:
This plugin is developes using mattermost plugin-2 architechture. 
Mattermost will switch plugin architecture in Mattermost v5.2 (will be released in 2018/8/16).

So if now you use this plugin, you run mattermost from latest [master branch](https://github.com/mattermost/mattermost-server).


### Custom Slash Command
1. Clone this repository `git clone https://github.com/kaakaa/mattermost-emojigen`
2. Install dependencies `dep ensure` (Need [golang/dep](https://github.com/golang/dep))
3. Write config.json (ExampleFile: `.config.json`)
  * **listen**: Mattermost emojigen server
  * **mattermost_url**: Mattermost server URL
  * **access_token**: [Mattermost Personal Access Token](https://docs.mattermost.com/developer/personal-access-tokens.html) for creating emoji
4. Run `go run cmd/emojigen/slash.main.go`
5. Create custom slash command
  * **DisplayName**: emojigen
  * **Description**: Creating emoji by https://github.com/kaakaa/mattermost-emojigen
  * **Command Trigger Word**: emojigen
  * **Request URL**: http://localhost:8505/emoji
  * **Request Method**: POST
  * **Autocomplete**: Checked
  * **Autocomplete Hint**: [EMOJI_NAME] [TEXT]
  * **Autocomplete Description**: Creating Emoji

## Usage

1. Executing `emojigen` command
```
/emojigen yabasugi やばすぎ
```
2. Use emoji (e.g.: `:yabasugi:`)

(For Japanese) 平仮名は絵文字にできますが漢字はほんの一部のみ画像化できるようです（カタカナはダメ）。https://github.com/pbnjay/pixfont で漢字なども画像化できる方法があれば可能になるかもしれません。

## Development

### Prerequires

Since `mattermost-emojigen` uses [go\-task/task](https://github.com/go-task/task) for effective development, you are better to install it.

```
go get -u github.com/go-task/task/cmd/task
```

### Building

```
task dist
```

# License

* MIT
  * see [LICENSE](LICENSE)