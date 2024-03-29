package font

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

// EmojiInfo is the struct that store emoji info by users
type EmojiInfo struct {
	Name            string
	Text            string
	FontColor       Color
	BackgroundColor Color
}

// Color is type of colors
type Color int

const (
	// Black represents black color
	Black Color = iota
	// Red represents red color
	Red
	// Green represents green color
	Green
	// Blue represents blue color
	Blue
	// White represents white color
	White
)

// ColorFromString convert string to Color type
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
	return Black, fmt.Errorf("invalid color name: [%s]", c)
}

// RGBA convert Color type to image.Uniform type
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

// NewEmojiInfoFromLine parse text into EmojiInfo
func NewEmojiInfoFromLine(text string) (*EmojiInfo, error) {
	args := strings.Split(text, " ")
	if len(args) < 3 || 5 < len(args) {
		return nil, fmt.Errorf("invalid command")
	}
	emoji := &EmojiInfo{}
	emoji.Name = args[1]
	emoji.Text = args[2]
	switch len(args) {
	case 4:
		c, err := ColorFromString(args[3])
		if err != nil {
			return nil, err
		}
		emoji.FontColor = c
		emoji.BackgroundColor = White
	case 5:
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
	default:
		emoji.FontColor = Black
		emoji.BackgroundColor = White
	}

	return emoji, nil
}
