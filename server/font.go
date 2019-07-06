package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/pbnjay/pixfont"
)

func generate(msg string) ([]byte, error) {
	r := []rune(msg)
	switch len(r) {
	case 0:
		return nil, fmt.Errorf("emoji text must container between 1 and 4 characters")
	case 1, 2:
		return genOneLineImage(r)
	case 3, 4:
		return genMultiLineImage(r)
	default:
		return nil, fmt.Errorf("emoji text must be less than or equal to 4 characters")
	}
}

func genOneLineImage(msg []rune) ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 8))
	pixfont.DrawString(img, 0, 0, string(msg), color.Black)

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func genMultiLineImage(msg []rune) ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 16))
	pixfont.DrawString(img, 0, 0, string(msg[0:2]), color.Black)
	pixfont.DrawString(img, 0, 8, string(msg[2:]), color.Black)

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
