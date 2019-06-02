package font

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
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
	fontSize = 64
	Width    = 128
	Height   = 128
)

type EmojiDrawer struct {
	baseDrawer font.Drawer
}

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
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		},
	}, nil
}

func (e *EmojiDrawer) GenerateEmoji(text string) ([]byte, error) {
	// TODO: must prevent to run concurrently
	switch utf8.RuneCountInString(text) {
	case 1, 2:
		return e.generateOneLineEmoji(text)
	case 3, 4:
		return e.generateTwoLinesEmoji(text)
	default:
		return nil, fmt.Errorf("Emojigen can generate 1~4 characters. You specified %d characters.", utf8.RuneCountInString(text))
	}
}

func (e *EmojiDrawer) generateOneLineEmoji(text string) ([]byte, error) {
	// Generate new Image
	e.baseDrawer.Dst = getNewImage()
	// To wide for fitting margin
	text = width.Widen.String(text)

	drawer := e.baseDrawer
	drawer.Dot.X = (fixed.I(Width) - (&drawer).MeasureString(text)) / 2
	drawer.Dot.Y = fixed.I((fontSize * 2 / 5) + (Height / 2))

	drawer.DrawString(text)
	buf := &bytes.Buffer{}
	if err := png.Encode(buf, drawer.Dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *EmojiDrawer) generateTwoLinesEmoji(text string) ([]byte, error) {
	// Generate new Image
	e.baseDrawer.Dst = getNewImage()

	// To wide for fitting margin
	r := []rune(text)
	t1 := width.Widen.String(string(r[0:2]))
	t2 := width.Widen.String(string(r[2:]))
	if utf8.RuneCountInString(t2) == 1 {
		t2 = t2 + width.Widen.String(" ")
	}

	// Calcurating positions X/Y is curious...
	drawer := e.baseDrawer
	drawer.Dot.X = (fixed.I(Width) - (&drawer).MeasureString(t1)) / 2
	drawer.Dot.Y = fixed.I(fontSize - fontSize/8)
	drawer.DrawString(t1)

	drawer.Dot.X = (fixed.I(Width) - (&drawer).MeasureString(t2)) / 2
	drawer.Dot.Y = fixed.I(Height - fontSize/8)
	drawer.DrawString(t2)

	buf := &bytes.Buffer{}
	if err := png.Encode(buf, drawer.Dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getNewImage() draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			img.Set(w, h, color.White)
		}
	}
	return img
}
