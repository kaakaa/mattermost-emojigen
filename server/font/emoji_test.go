package font

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmojiInfoFromLine(t *testing.T) {
	for name, test := range map[string]struct {
		text        string
		shouldError bool
		expected    *EmojiInfo
	}{
		"Simple command": {
			text:        "/emojigen emoji_name emoji_text",
			shouldError: false,
			expected:    &EmojiInfo{Name: "emoji_name", Text: "emoji_text", FontColor: Black, BackgroundColor: White},
		},
		"Command with FontColor": {
			text:        "/emojigen emoji_name emoji_text Red",
			shouldError: false,
			expected:    &EmojiInfo{Name: "emoji_name", Text: "emoji_text", FontColor: Red, BackgroundColor: White},
		},
		"Command with FontColor and BackgroundColor": {
			text:        "/emojigen emoji_name emoji_text Red Black",
			shouldError: false,
			expected:    &EmojiInfo{Name: "emoji_name", Text: "emoji_text", FontColor: Red, BackgroundColor: Black},
		},
		"Various capitelization": {
			text:        "/emojigen emoji_name emoji_text BLUE green",
			shouldError: false,
			expected:    &EmojiInfo{Name: "emoji_name", Text: "emoji_text", FontColor: Blue, BackgroundColor: Green},
		},
		"Error, invalid color": {
			text:        "/emojigen emoji_name emoji_text INVALID COLOR",
			shouldError: true,
		},
		"Error, invalid command1": {
			text:        "/emojigen emoji_name",
			shouldError: true,
		},
		"Error, invalid command2": {
			text:        "/emojigen emoji_name emoji_text Black White INVALID",
			shouldError: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := NewEmojiInfoFromLine(test.text)
			if test.shouldError {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(test.expected, actual)
			}
		})
	}
}
