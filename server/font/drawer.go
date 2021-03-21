package font

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"golang.org/x/text/width"
)

const (
	imageSize = 128
)

// EmojiDrawer is the struct for drawing emoji
type EmojiDrawer struct {
	tf *truetype.Font
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
	return &EmojiDrawer{
		tf: tf,
	}, nil
}

// GenerateEmoji generate the image of emoji and returns it as []byte format
func (e *EmojiDrawer) GenerateEmoji(emoji *EmojiInfo) ([]byte, error) {
	lines := strings.Split(emoji.Text, "\n")

	maxLen := len(lines)
	maxLineLen := 0
	runes := [][]rune{}
	for _, l := range lines {
		r := []rune(l)
		runes = append(runes, r)
		if maxLen < len(r) {
			maxLen = len(r)
		}
		if maxLineLen < len(r) {
			maxLineLen = len(r)
		}
	}

	// Generate new Image
	drawer := e.getDrawer(maxLen)
	drawer.Dst = getNewImage(emoji.BackgroundColor.RGBA())
	drawer.Src = emoji.FontColor.RGBA()

	unitSize := (&drawer).MeasureString(width.Widen.String("0"))
	yOrigin := fixed.Int26_6((fixed.I(imageSize)-unitSize*fixed.Int26_6(len(runes)))/2) - fixed.Int26_6(unitSize/8)
	for i, r := range runes {
		t := width.Widen.String(string(r))
		if utf8.RuneCountInString(t) < maxLineLen {
			t = t + width.Widen.String(strings.Repeat(" ", maxLineLen-utf8.RuneCountInString(t)))
		}

		drawer.Dot.X = (fixed.I(imageSize) - (&drawer).MeasureString(t)) / 2
		drawer.Dot.Y = yOrigin + (fixed.Int26_6(i+1) * unitSize)
		drawer.DrawString(t)
	}

	buf := &bytes.Buffer{}
	if err := png.Encode(buf, drawer.Dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *EmojiDrawer) getDrawer(len int) font.Drawer {
	size := imageSize / len
	opts := truetype.Options{
		Size:              float64(size),
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}
	face := truetype.NewFace(e.tf, &opts)
	return font.Drawer{
		Src:  image.White,
		Face: face,
		Dot:  fixed.Point26_6{},
	}
}
func getNewImage(color *image.Uniform) draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	for h := 0; h < imageSize; h++ {
		for w := 0; w < imageSize; w++ {
			img.Set(w, h, color)
		}
	}
	return img
}
