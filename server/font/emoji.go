package font

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

type EmojiInfo struct {
	Name            string
	Text            string
	FontColor       Color
	BackgroundColor Color
}

type Color int

const (
	Black Color = iota
	Red
	Green
	Blue
	White
)

func ColorFromString(c string) (Color, error) {
	switch strings.ToLower(c) {
	case "black":
		return Black, nil
	case "red":
		return Red, nil
	case "green":
		return Green, nil
	case "blue":
		return Blue, nil
	case "white":
		return White, nil
	}
	return Black, fmt.Errorf("Invalid color name: [%s]", c)
}

func (c Color) RGBA() *image.Uniform {
	switch c {
	case Black:
		return image.Black
	case Red:
		return image.NewUniform(color.RGBA{255, 0, 0, 255})
	case Green:
		return image.NewUniform(color.RGBA{0, 255, 0, 255})
	case Blue:
		return image.NewUniform(color.RGBA{0, 0, 255, 255})
	case White:
		return image.White
	}
	return image.Black
}

func NewEmojiInfoFromLine(text string) (*EmojiInfo, error) {
	args := strings.Split(text, " ")
	if len(args) < 3 || 5 < len(args) {
		return nil, fmt.Errorf("Invalid command")
	}
	emoji := &EmojiInfo{}
	emoji.Name = args[1]
	emoji.Text = args[2]
	if len(args) == 4 {
		c, err := ColorFromString(args[3])
		if err != nil {
			return nil, err
		}
		emoji.FontColor = c
		emoji.BackgroundColor = White
	} else if len(args) == 5 {
		c, err := ColorFromString(args[3])
		if err != nil {
			return nil, err
		}
		emoji.FontColor = c
		c, err = ColorFromString(args[4])
		if err != nil {
			return nil, err
		}
		emoji.BackgroundColor = c
	} else {
		emoji.FontColor = Black
		emoji.BackgroundColor = White
	}

	return emoji, nil
}
