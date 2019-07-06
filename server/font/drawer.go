package font

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"path/filepath"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"golang.org/x/text/width"
)

const (
	fontSize    = 64
	imageWidth  = 128
	imageHeight = 128
)

// EmojiDrawer is the struct for drawing emoji
type EmojiDrawer struct {
	baseDrawer font.Drawer
}

// NewEmojiDrawer returns initialized EmojiDrawer
func NewEmojiDrawer(bundlePath string) (*EmojiDrawer, error) {
	path := filepath.Join(bundlePath, "assets", "ttf", "mplus", "mplus-2p-black.ttf")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tf, err := truetype.Parse(b)
	if err != nil {
		return nil, err
	}
	opts := truetype.Options{
		Size:              fontSize,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}
	face := truetype.NewFace(tf, &opts)

	return &EmojiDrawer{
		baseDrawer: font.Drawer{
			Src:  image.White,
			Face: face,
			Dot:  fixed.Point26_6{},
		},
	}, nil
}

// GenerateEmoji generate the image of moji and returns it as []byte format
func (e *EmojiDrawer) GenerateEmoji(emoji *EmojiInfo) ([]byte, error) {
	// TODO: must prevent to run concurrently
	switch utf8.RuneCountInString(emoji.Text) {
	case 1, 2:
		return e.generateOneLineEmoji(emoji)
	case 3, 4:
		return e.generateTwoLinesEmoji(emoji)
	default:
		return nil, fmt.Errorf("emojigen can generate 1~4 characters. You specified %d characters", utf8.RuneCountInString(emoji.Text))
	}
}

func (e *EmojiDrawer) generateOneLineEmoji(emoji *EmojiInfo) ([]byte, error) {
	// Generate new Image
	e.baseDrawer.Dst = getNewImage(emoji.BackgroundColor.RGBA())
	// To wide for fitting margin
	text := width.Widen.String(emoji.Text)

	drawer := e.baseDrawer
	drawer.Src = emoji.FontColor.RGBA()

	drawer.Dot.X = (fixed.I(imageWidth) - (&drawer).MeasureString(text)) / 2
	drawer.Dot.Y = fixed.I((fontSize * 2 / 5) + (imageHeight / 2))

	drawer.DrawString(text)
	buf := &bytes.Buffer{}
	if err := png.Encode(buf, drawer.Dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *EmojiDrawer) generateTwoLinesEmoji(emoji *EmojiInfo) ([]byte, error) {
	// Generate new Image
	e.baseDrawer.Dst = getNewImage(emoji.BackgroundColor.RGBA())

	// To wide for fitting margin
	r := []rune(emoji.Text)
	t1 := width.Widen.String(string(r[0:2]))
	t2 := width.Widen.String(string(r[2:]))
	if utf8.RuneCountInString(t2) == 1 {
		t2 = t2 + width.Widen.String(" ")
	}

	// Calcurating positions X/Y is curious...
	drawer := e.baseDrawer
	drawer.Src = emoji.FontColor.RGBA()

	drawer.Dot.X = (fixed.I(imageWidth) - (&drawer).MeasureString(t1)) / 2
	drawer.Dot.Y = fixed.I(fontSize - fontSize/8)
	drawer.DrawString(t1)

	drawer.Dot.X = (fixed.I(imageWidth) - (&drawer).MeasureString(t2)) / 2
	drawer.Dot.Y = fixed.I(imageHeight - fontSize/8)
	drawer.DrawString(t2)

	buf := &bytes.Buffer{}
	if err := png.Encode(buf, drawer.Dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getNewImage(color *image.Uniform) draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for h := 0; h < imageHeight; h++ {
		for w := 0; w < imageWidth; w++ {
			img.Set(w, h, color)
		}
	}
	return img
}
