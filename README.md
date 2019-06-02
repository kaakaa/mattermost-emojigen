[![Build Status](https://travis-ci.org/kaakaa/mattermost-emojigen.svg?branch=master)](https://travis-ci.org/kaakaa/mattermost-emojigen)

# Mattermost EmojiGen

Mattermost plugin for generating custom emoji.

![sample](./emoji_sample.png)

## Set up

### Plugin
1. Download a plugin distribution from [Releases · kaakaa/mattermost\-emojigen](https://github.com/kaakaa/mattermost-emojigen/releases)
2. Upload and Enabling plugin from your mattermost's admin console
3. Set your Mattermost Personal Access Token in configuration page on admin console

## Usage

1. Executing `emojigen` command
```
/emojigen yabasugi やばすぎ
```
2. Use emoji (e.g.: `:yabasugi:`)

## Development

### Building

```
make dist
```

# License

* This plugin is distributed under [MIT LICENSE](LICENSE)
* This plugin uses [**M+ Fonts**](https://mplus-fonts.osdn.jp/) for generating emojis. **M+ Fonts** is distributed under [LICENSE](./assets/ttf/mplus/LICENSE_E).