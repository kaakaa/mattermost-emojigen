package font

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateEmoji(t *testing.T) {
	drawer, _ := NewEmojiDrawer("../../")
	for name, test := range map[string]struct {
		sut         *EmojiDrawer
		emoji       *EmojiInfo
		shouldError bool
		byteLength  int
	}{
		"Color_1":           {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "字赤背白", FontColor: Red, BackgroundColor: White}, shouldError: false, byteLength: 2386},
		"Color_2":           {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "字黒背青", FontColor: Black, BackgroundColor: Blue}, shouldError: false, byteLength: 1948},
		"Color_3":           {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "字緑背黒", FontColor: Green, BackgroundColor: Black}, shouldError: false, byteLength: 2397},
		"EmptyLine":         {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "AB\nCD", FontColor: Green, BackgroundColor: Black}, shouldError: false, byteLength: 2652},
		"Hiraganas_3":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ぱぴぷ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3818},
		"Hiraganas_4":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "やばすぎ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3619},
		"Katakanas_3":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ギャグ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2828},
		"Katakanas_4":       {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ピョンプ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2397},
		"Kanjis_3":          {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "神奈川", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2701},
		"Kanjis_4":          {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "四面楚歌", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3043},
		"Kanjis_4_2lines":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "四面\n楚歌", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 4640},
		"Kanjis_4_2lines_2": {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "四面楚\n歌", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3695},
		"Kanjis_4_3lines":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "四面\n楚\n歌", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3600},
		"Kanjis_4_4lines":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "四\n面\n楚\n歌", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 3161},
		"Kanjis_9_3lines":   {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "圧倒的\n当事者\n意識", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 6607},
		"Kanjis_9_3lines_2": {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "圧倒的\n当事者\n　意識", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 6575},
		"Lower_case_1":      {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "a", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2157},
		"Lower_case_2":      {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "ab", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2026},
		"Lower_case_3":      {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "abc", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1962},
		"Symbols_3":         {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "(!)", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1760},
		"Symbols_4":         {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "[~/]", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1409},
		"Upper_cases_3":     {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "XYZ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 1916},
		"Upper_cases_4":     {sut: drawer, emoji: &EmojiInfo{Name: "name", Text: "WXYZ", FontColor: Black, BackgroundColor: White}, shouldError: false, byteLength: 2298},
	} {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			b, err := test.sut.GenerateEmoji(test.emoji)
			if test.shouldError {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(test.byteLength, len(b))
				// for debug
				// ioutil.WriteFile(fmt.Sprintf("./TestGenerateEmoji_%s.png", name), b, os.ModePerm)
			}
		})
	}
}
