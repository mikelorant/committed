package emoji

import (
	_ "embed"
	"strings"
)

//go:embed emoji.txt
var mockEmoji string

func Emoji() string {
	return strings.TrimSpace(mockEmoji)
}
