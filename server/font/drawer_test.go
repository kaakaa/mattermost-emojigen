package font

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOneLineEmoji(t *testing.T) {
	drawer, _ := NewEmojiDrawer("../../")
	for name, test := range map[string]struct {
		sut         *EmojiDrawer
		text        string
		shouldError bool
		byteLength  int
	}{
		"Lower_case_1":  {sut: drawer, text: "a", shouldError: false, byteLength: 1368},
		"Lower_cases_2": {sut: drawer, text: "ab", shouldError: false, byteLength: 2026},
		"Upper_case_1":  {sut: drawer, text: "Z", shouldError: false, byteLength: 956},
		"Upper_case2_2": {sut: drawer, text: "YZ", shouldError: false, byteLength: 1616},
		"Hiragana_1":    {sut: drawer, text: "あ", shouldError: false, byteLength: 1978},
		"Hiraganas_2":   {sut: drawer, text: "あい", shouldError: false, byteLength: 3222},
		"Katakana_1":    {sut: drawer, text: "ン", shouldError: false, byteLength: 1647},
		"Katakanas_2":   {sut: drawer, text: "ヲン", shouldError: false, byteLength: 2279},
		"Knaji_1":       {sut: drawer, text: "蕾", shouldError: false, byteLength: 1263},
		"Knajis_2":      {sut: drawer, text: "令和", shouldError: false, byteLength: 2619},
		"Symbol_1":      {sut: drawer, text: "*", shouldError: false, byteLength: 1556},
		"Symbols_2":     {sut: drawer, text: "!?", shouldError: false, byteLength: 1412},
	} {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			test.sut.baseDrawer.Dst = getNewImage()
			b, err := test.sut.generateOneLineEmoji(test.text)
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
		text        string
		shouldError bool
		byteLength  int
	}{
		"Lower_cases_3": {sut: drawer, text: "abc", shouldError: false, byteLength: 2780},
		"Lower_cases_4": {sut: drawer, text: "abcd", shouldError: false, byteLength: 3312},
		"Upper_cases_3": {sut: drawer, text: "XYZ", shouldError: false, byteLength: 2511},
		"Upper_cases_4": {sut: drawer, text: "WXYZ", shouldError: false, byteLength: 3753},
		"Hiraganas_3":   {sut: drawer, text: "ぱぴぷ", shouldError: false, byteLength: 5607},
		"Hiraganas_4":   {sut: drawer, text: "やばすぎ", shouldError: false, byteLength: 6360},
		"Katakanas_3":   {sut: drawer, text: "ギャグ", shouldError: false, byteLength: 3853},
		"Katakanas_4":   {sut: drawer, text: "ピョンプ", shouldError: false, byteLength: 3828},
		"Knajis_3":      {sut: drawer, text: "神奈川", shouldError: false, byteLength: 3503},
		"Knajis_4":      {sut: drawer, text: "四面楚歌", shouldError: false, byteLength: 4640},
		"Symbols_3":     {sut: drawer, text: "(!)", shouldError: false, byteLength: 2465},
		"Symbols_4":     {sut: drawer, text: "[~/]", shouldError: false, byteLength: 1696},
	} {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			test.sut.baseDrawer.Dst = getNewImage()
			b, err := test.sut.generateTwoLinesEmoji(test.text)
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
