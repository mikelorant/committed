package commit

import (
	"fmt"
	"strings"

	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
)

func MessageToEmoji(msg string) emoji.NullEmoji {
	ls := strings.Split(msg, "\n")
	fw := strings.Split(ls[0], " ")[0]

	return emoji.New().Find(fw)
}

func MessageToSummary(msg string) string {
	lines := strings.Split(msg, "\n")
	line := lines[0]

	if !hasSummary(msg) {
		return ""
	}

	ls := strings.Split(line, " ")
	fw := ls[0]

	if emoji.Has(fw) {
		if len(ls) <= 1 {
			return ""
		}
		line = strings.Join(ls[1:], " ")
	}

	return line
}

func MessageToBody(msg string) string {
	if !hasSummary(msg) {
		return msg
	}

	ls := strings.Split(msg, "\n")

	switch len(ls) {
	case 1:
		return ""
	case 2:
		return ls[1]
	}

	return strings.Join(ls[2:], "\n")
}

func EmojiSummaryToSubject(emoji, summary string) string {
	var subject string

	if emoji != "" {
		subject = fmt.Sprintf("%s %s", emoji, summary)
	} else {
		subject = summary
	}

	return subject
}

func UserToAuthor(user repository.User) string {
	if user.Name == "" || user.Email == "" {
		return ""
	}

	return fmt.Sprintf("%s <%s>", user.Name, user.Email)
}

func hasSummary(msg string) bool {
	ls := strings.Split(msg, "\n")

	switch len(ls) {
	case 1:
		if len(ls[0]) == 0 || len(ls[0]) > 72 {
			return false
		}
		return true
	case 2:
		if ls[0] != "" && ls[1] == "" {
			return true
		}
	default:
		if ls[0] != "" && ls[1] == "" && ls[2] != "" {
			return true
		}
	}

	return false
}
