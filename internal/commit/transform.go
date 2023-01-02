package commit

import (
	"fmt"
	"strings"

	"github.com/mikelorant/committed/internal/emoji"
)

func (c *Config) MessageToEmoji() emoji.NullEmoji {
	ls := strings.Split(c.Repository.Head.Message, "\n")
	fw := strings.Split(ls[0], " ")[0]

	return emoji.New().Find(fw)
}

func (c *Config) MessageToSummary() string {
	lines := strings.Split(c.Repository.Head.Message, "\n")
	line := lines[0]

	if !hasSummary(c.Repository.Head.Message) {
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

func (c *Config) MessageToBody() string {
	lines := c.Repository.Head.Message

	if !hasSummary(lines) {
		return lines
	}

	ls := strings.Split(c.Repository.Head.Message, "\n")

	switch len(ls) {
	case 1:
		return ""
	case 2:
		return ls[1]
	}

	return strings.Join(ls[2:], "\n")
}

func (c *Commit) EmojiSummaryToSubject() string {
	var subject string

	if c.Emoji != "" {
		subject = fmt.Sprintf("%s %s", c.Emoji, c.Summary)
	} else {
		subject = c.Summary
	}

	return subject
}

func (c *Commit) UserToAuthor() string {
	if c.Author.Name == "" || c.Author.Email == "" {
		return ""
	}

	return fmt.Sprintf("%s <%s>", c.Author.Name, c.Author.Email)
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
