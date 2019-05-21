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

(For Japanese) 平仮名は絵文字にできますが漢字はほんの一部のみ画像化できるようです（カタカナはダメ）。https://github.com/pbnjay/pixfont で漢字なども画像化できる方法があれば可能になるかもしれません。

## Development

### Building

```
make dist
```

# License

* MIT
  * see [LICENSE](LICENSE)
