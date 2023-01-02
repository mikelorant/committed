package commit

import (
	"strings"

	"github.com/mikelorant/committed/internal/emoji"
)

func messageToEmoji(msg string) emoji.NullEmoji {
	ls := strings.Split(msg, "\n")
	fw := strings.Split(ls[0], " ")[0]

	return emoji.New().Find(fw)
}

func messageToSummary(msg string) string {
	lines := strings.Split(msg, "\n")
	line := lines[0]
	ls := strings.Split(line, " ")
	fw := ls[0]

	if emoji.Has(fw) {
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
