package commit

import (
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

func (c *Config) MessageToBody() string {
	if hasSummary(c.Repository.Head.Message) {
		ls := strings.Split(c.Repository.Head.Message, "\n")
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
