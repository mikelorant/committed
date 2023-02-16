package commit

import (
	_ "embed"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/snapshot"
	"github.com/mikelorant/committed/internal/theme"
)

type State struct {
	Placeholders Placeholders
	Repository   repository.Description
	Emojis       *emoji.Set
	Theme        theme.Theme
	Config       config.Config
	Snapshot     snapshot.Snapshot
	Options      Options
	Hook         Hook
}

type Placeholders struct {
	Hash    string
	Summary string
	Body    string
	Help    string
}

type Config struct {
	View    config.View
	Commit  config.Commit
	Authors []repository.User
}

type Hook struct {
	Amend   bool
	Message string
}

//go:embed message.txt
var PlaceholderMessage string

//go:embed help.txt
var PlaceholderHelp string

const (
	PlaceholderHash    string = "1234567890abcdef1234567890abcdef12345678"
	PlaceholderSummary string = "Capitalized, short (50 chars or less) summary"
)

func placeholders() Placeholders {
	return Placeholders{
		Hash:    PlaceholderHash,
		Summary: PlaceholderSummary,
		Body:    PlaceholderMessage,
		Help:    PlaceholderHelp,
	}
}
