package commit

import (
	"strings"

	"github.com/mikelorant/committed/internal/emoji"
)

func messageToEmoji(msg string, es []emoji.Emoji) emoji.Emoji {
	ls := strings.Split(msg, "\n")
	fw := strings.Split(ls[0], " ")[0]

	if emoji.HasEmoji(fw) {
		for _, e := range es {
			if e.Character == fw {
				return emoji.Emoji{
					Character: fw,
					ShortCode: e.ShortCode,
				}
			}
		}
	}

	if !emoji.HasShortcode(fw) {
		return emoji.Emoji{}
	}

	for _, e := range es {
		if e.ShortCode == fw {
			return emoji.Emoji{
				Character: e.Character,
				ShortCode: fw,
			}
		}
	}

	return emoji.Emoji{}
}

func messageToSummary(msg string) string {
	lines := strings.Split(msg, "\n")
	line := lines[0]
	ls := strings.Split(line, " ")
	fw := ls[0]

	if emoji.HasEmoji(fw) || emoji.HasShortcode(fw) {
		if len(line) <= 1 {
			return ""
		}
		return strings.Join(ls[1:], " ")
	}

	return line
}

func messageToBody(msg string) string {
	if hasSummary(msg) {
		ls := strings.Split(msg, "\n")
		return strings.Join(ls[2:], "\n")
	}

	return ""
}

func hasSummary(msg string) bool {
	ls := strings.Split(msg, "\n")
	if len(ls) <= 2 {
		return false
	}

	if ls[0] != "" && ls[1] == "" && ls[2] != "" {
		return true
	}

	return false
}
