package commit

import (
	_ "embed"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
)

type State struct {
	Placeholders Placeholders
	Repository   repository.Description
	Emojis       []emoji.Emoji
	Config       config.Config
	Options      Options
}

type Placeholders struct {
	Hash    string
	Summary string
	Body    string
}

type Config struct {
	View    config.View
	Commit  config.Commit
	Authors []repository.User
}

//go:embed message.txt
var PlaceholderMessage string

const (
	PlaceholderHash    string = "1234567890abcdef1234567890abcdef1234567890"
	PlaceholderSummary string = "Capitalized, short (50 chars or less) summary"
)

func placeholders() Placeholders {
	return Placeholders{
		Hash:    PlaceholderHash,
		Summary: PlaceholderSummary,
		Body:    PlaceholderMessage,
	}
}
