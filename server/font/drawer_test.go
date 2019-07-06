package font

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOneLineEmoji(t *testing.T) {
	drawer, _ := NewEmojiDrawer("../../")
	for name, test := range map[string]struct {
		sut         *EmojiDrawer
		emoji       *EmojiInfo
		shouldError bool
		byteLength  int
	}{
		"Lower_case_1":  {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "a", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1368},
		"Lower_cases_2": {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ab", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2026},
		"Upper_case_1":  {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "Z", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 956},
		"Upper_case2_2": {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "YZ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1616},
		"Hiragana_1":    {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "あ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1978},
		"Hiraganas_2":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "あい", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3222},
		"Katakana_1":    {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ン", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1647},
		"Katakanas_2":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ヲン", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2279},
		"Knaji_1":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "蕾", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1263},
		"Knajis_2":      {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "令和", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2619},
		"Symbol_1":      {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "*", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1556},
		"Symbols_2":     {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "!?", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1412},
		"Color_1":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "赤", FontColor: Red, BackgroundColor: White}, shouldError: false, byteLength: 1849},
		"Color_2":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "青", FontColor: Black, BackgroundColor: Blue}, shouldError: false, byteLength: 906},
		"Color_3":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "逆", FontColor: White, BackgroundColor: Black}, shouldError: false, byteLength: 1924},
	} {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			b, err := test.sut.generateOneLineEmoji(test.emoji)
			if test.shouldError {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(test.byteLength, len(b))
				// for debug
				// ioutil.WriteFile(fmt.Sprintf("./testOne_%s.png", name), b, os.ModePerm)
			}
		})
	}
}

func TestGenerateTwoLinesEmoji(t *testing.T) {
	drawer, _ := NewEmojiDrawer("../../")
	for name, test := range map[string]struct {
		sut         *EmojiDrawer
		emoji       *EmojiInfo
		shouldError bool
		byteLength  int
	}{
		"Lower_cases_3": {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "abc", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2780},
		"Lower_cases_4": {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "abcd", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3312},
		"Upper_cases_3": {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "XYZ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2511},
		"Upper_cases_4": {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "WXYZ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3753},
		"Hiraganas_3":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ぱぴぷ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 5607},
		"Hiraganas_4":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "やばすぎ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 6360},
		"Katakanas_3":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ギャグ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3853},
		"Katakanas_4":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ピョンプ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3828},
		"Knajis_3":      {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "神奈川", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3503},
		"Knajis_4":      {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "四面楚歌", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 4640},
		"Symbols_3":     {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "(!)", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2465},
		"Symbols_4":     {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "[~/]", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1696},
		"Color_1":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "字赤背白", FontColor: Red, BackgroundColor: White}, shouldError: false, byteLength: 3475},
		"Color_2":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "字黒背青", FontColor: Black, BackgroundColor: Blue}, shouldError: false, byteLength: 2905},
		"Color_3":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "字緑背黒", FontColor: Green, BackgroundColor: Black}, shouldError: false, byteLength: 5405},
	} {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			b, err := test.sut.generateTwoLinesEmoji(test.emoji)
			if test.shouldError {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(test.byteLength, len(b))
				// for debug
				// ioutil.WriteFile(fmt.Sprintf("./testTwo_%s.png", name), b, os.ModePerm)
			}
		})
	}
}
