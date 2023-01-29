package theme_test

import (
	"fmt"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/ui/theme"

	"github.com/stretchr/testify/assert"
)

type Colour struct {
	Dark  string
	Light string
}

type body struct {
	Boundary            Colour
	TextAreaPlaceholder Colour
	TextAreaPrompt      Colour
	TextAreaFocusedText Colour
	TextAreaBlurredText Colour
	TextAreaCursorStyle Colour
}

type filterlist struct {
	Boundary                  Colour
	ListNormalTitle           Colour
	ListSelectedTitle         Colour
	ListNoItems               Colour
	TextInputPromptMark       Colour
	TextInputPromptText       Colour
	PaginatorDots             Colour
	TextInputPromptStyle      Colour
	TextInputTextStyle        Colour
	TextInputPlaceholderStyle Colour
	TextInputCursorStyle      Colour
}

type footer struct {
	View Colour
}

type header struct {
	EmojiBoundary                Colour
	SummaryBoundary              Colour
	CounterDivider               Colour
	CounterLimit                 Colour
	SummaryInputPromptStyle      Colour
	SummaryInputTextStyle        Colour
	SummaryInputPlaceholderStyle Colour
	SummaryInputCursorStyle      Colour
	CounterDefault               Colour
	CounterLow                   Colour
	CounterNormal                Colour
	CounterWarning               Colour
	CounterHigh                  Colour
	ReadyError                   Colour
	ReadyIncomplete              Colour
	ReadyOK                      Colour
	CommitTypeNew                Colour
	CommitTypeAmend              Colour
}

type help struct {
	Boundary Colour
	Viewport Colour
}

type info struct {
	HashText            Colour
	HashValue           Colour
	BranchHead          Colour
	BranchLocal         Colour
	BranchGrouping      Colour
	BranchRemote        Colour
	Colon               Colour
	AuthorAngledBracket Colour
	AuthorText          Colour
	AuthorValue         Colour
	DateText            Colour
	DateValue           Colour
}

type message struct {
	Summary Colour
	Body    Colour
	Footer  Colour
}

type shortcut struct {
	Key          Colour
	Label        Colour
	Plus         Colour
	AngleBracket Colour
}

func TestBody(t *testing.T) {
	tests := []struct {
		name string
		body body
	}{
		{
			name: "body",
			body: body{
				Boundary:            Colour{Dark: "#bbbbbb"},
				TextAreaPlaceholder: Colour{Dark: "#555555", Light: "#555555"},
				TextAreaPrompt:      Colour{Dark: "#bbbbbb"},
				TextAreaFocusedText: Colour{Dark: "#bbbbbb"},
				TextAreaBlurredText: Colour{Dark: "#bbbbbb"},
				TextAreaCursorStyle: Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)
			body := th.Body()

			assert.Equal(t, tt.body.Boundary, toColour(body.Boundary), "Boundary")
			assert.Equal(t, tt.body.TextAreaPlaceholder, toColour(body.TextAreaPlaceholder), "TextAreaPlaceholder")
			assert.Equal(t, tt.body.TextAreaPrompt, toColour(body.TextAreaPrompt), "TextAreaPrompt")
			assert.Equal(t, tt.body.TextAreaFocusedText, toColour(body.TextAreaFocusedText), "TextAreaFocusedText")
			assert.Equal(t, tt.body.TextAreaBlurredText, toColour(body.TextAreaBlurredText), "TextAreaBlurredText")
			assert.Equal(t, tt.body.TextAreaCursorStyle, toColour(body.TextAreaCursorStyle), "TextAreaCursorStyle")
		})
	}
}

func TestFilterList(t *testing.T) {
	tests := []struct {
		name       string
		filterlist filterlist
	}{
		{
			name: "filterlist",
			filterlist: filterlist{
				Boundary:                  Colour{Dark: "#bbbbbb"},
				ListNormalTitle:           Colour{Dark: "#bbbbbb"},
				ListSelectedTitle:         Colour{Dark: "#00bbbb", Light: "#bb0000"},
				ListNoItems:               Colour{Dark: "#555555", Light: "#555555"},
				TextInputPromptMark:       Colour{Dark: "#00bb00", Light: "#bb00bb"},
				TextInputPromptText:       Colour{Dark: "#bbbbbb"},
				PaginatorDots:             Colour{Dark: "#00bbbb", Light: "#bb0000"},
				TextInputPromptStyle:      Colour{Dark: "#bbbbbb"},
				TextInputTextStyle:        Colour{Dark: "#bbbbbb"},
				TextInputPlaceholderStyle: Colour{Dark: "#555555", Light: "#555555"},
				TextInputCursorStyle:      Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)
			filterlist := th.FilterList()

			assert.Equal(t, tt.filterlist.Boundary, toColour(filterlist.Boundary), "Boundary")
			assert.Equal(t, tt.filterlist.ListNormalTitle, toColour(filterlist.ListNormalTitle), "ListNormalTitle")
			assert.Equal(t, tt.filterlist.ListSelectedTitle, toColour(filterlist.ListSelectedTitle), "ListSelectedTitle")
			assert.Equal(t, tt.filterlist.ListNoItems, toColour(filterlist.ListNoItems), "ListNoItems")
			assert.Equal(t, tt.filterlist.TextInputPromptMark, toColour(filterlist.TextInputPromptMark), "TextInputPromptMark")
			assert.Equal(t, tt.filterlist.TextInputPromptText, toColour(filterlist.TextInputPromptText), "TextInputPromptText")
			assert.Equal(t, tt.filterlist.PaginatorDots, toColour(filterlist.PaginatorDots), "PaginatorDots")
			assert.Equal(t, tt.filterlist.TextInputPromptStyle, toColour(filterlist.TextInputPromptStyle), "TextInputPromptStyle")
			assert.Equal(t, tt.filterlist.TextInputTextStyle, toColour(filterlist.TextInputTextStyle), "TextInputTextStyle")
			assert.Equal(t, tt.filterlist.TextInputPlaceholderStyle, toColour(filterlist.TextInputPlaceholderStyle), "TextInputPlaceholderStyle")
			assert.Equal(t, tt.filterlist.TextInputCursorStyle, toColour(filterlist.TextInputCursorStyle), "TextInputCursorStyle")
		})
	}
}

func TestFooter(t *testing.T) {
	tests := []struct {
		name   string
		footer footer
	}{
		{
			name: "Footer",
			footer: footer{
				View: Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)
			footer := th.Footer()

			assert.Equal(t, tt.footer.View, toColour(footer.View), "Boundary")
		})
	}
}

func TestHeader(t *testing.T) {
	tests := []struct {
		name   string
		header header
	}{
		{
			name: "Header",
			header: header{
				EmojiBoundary:                Colour{Dark: "#bbbbbb"},
				SummaryBoundary:              Colour{Dark: "#bbbbbb"},
				CounterDivider:               Colour{Dark: "#bbbbbb"},
				CounterLimit:                 Colour{Dark: "#bbbbbb"},
				SummaryInputPromptStyle:      Colour{Dark: "#bbbbbb"},
				SummaryInputTextStyle:        Colour{Dark: "#bbbbbb"},
				SummaryInputPlaceholderStyle: Colour{Dark: "#555555", Light: "#555555"},
				SummaryInputCursorStyle:      Colour{Dark: "#bbbbbb"},
				CounterDefault:               Colour{Dark: "#bbbbbb"},
				CounterLow:                   Colour{Dark: "#bbbb00", Light: "#0000bb"},
				CounterNormal:                Colour{Dark: "#00bb00", Light: "#bb00bb"},
				CounterWarning:               Colour{Dark: "#bbbb00", Light: "#0000bb"},
				CounterHigh:                  Colour{Dark: "#ff5555", Light: "#55ffff"},
				ReadyError:                   Colour{Dark: "#ff5555", Light: "#55ffff"},
				ReadyIncomplete:              Colour{Dark: "#bbbb00", Light: "#0000bb"},
				ReadyOK:                      Colour{Dark: "#00bb00", Light: "#bb00bb"},
				CommitTypeNew:                Colour{Dark: "#00bb00", Light: "#bb00bb"},
				CommitTypeAmend:              Colour{Dark: "#bbbb00", Light: "#0000bb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)
			header := th.Header()

			assert.Equal(t, tt.header.EmojiBoundary, toColour(header.EmojiBoundary), "EmojiBoundary")
			assert.Equal(t, tt.header.SummaryBoundary, toColour(header.SummaryBoundary), "SummaryBoundary")
			assert.Equal(t, tt.header.CounterDivider, toColour(header.CounterDivider), "CounterDivider")
			assert.Equal(t, tt.header.CounterLimit, toColour(header.CounterLimit), "CounterLimit")
			assert.Equal(t, tt.header.SummaryInputPromptStyle, toColour(header.SummaryInputPromptStyle), "SummaryInputPromptStyle")
			assert.Equal(t, tt.header.SummaryInputTextStyle, toColour(header.SummaryInputTextStyle), "SummaryInputTextStyle")
			assert.Equal(t, tt.header.SummaryInputPlaceholderStyle, toColour(header.SummaryInputPlaceholderStyle), "SummaryInputPlaceholderStyle")
			assert.Equal(t, tt.header.SummaryInputCursorStyle, toColour(header.SummaryInputCursorStyle), "SummaryInputCursorStyle")
			assert.Equal(t, tt.header.CounterDefault, toColour(header.CounterDefault), "CounterDefault")
			assert.Equal(t, tt.header.CounterLow, toColour(header.CounterLow), "CounterLow")
			assert.Equal(t, tt.header.CounterNormal, toColour(header.CounterNormal), "CounterNormal")
			assert.Equal(t, tt.header.CounterWarning, toColour(header.CounterWarning), "CounterWarning")
			assert.Equal(t, tt.header.CounterHigh, toColour(header.CounterHigh), "CounterHigh")
			assert.Equal(t, tt.header.CommitTypeNew, toColour(header.CommitTypeNew), "CommitTypeNew")
			assert.Equal(t, tt.header.CommitTypeAmend, toColour(header.CommitTypeAmend), "CommitTypeAmend")
		})
	}
}

func TestHelp(t *testing.T) {
	tests := []struct {
		name string
		help help
	}{
		{
			name: "Help",
			help: help{
				Boundary: Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)
			help := th.Help()

			assert.Equal(t, tt.help.Boundary, toColour(help.Boundary), "Boundary")
		})
	}
}

func TestInfo(t *testing.T) {
	tests := []struct {
		name string
		info info
	}{
		{
			name: "Info",
			info: info{
				HashText:            Colour{Dark: "#bbbb00", Light: "#0000bb"},
				HashValue:           Colour{Dark: "#bbbb00", Light: "#0000bb"},
				BranchHead:          Colour{Dark: "#55ffff", Light: "#ff5555"},
				BranchLocal:         Colour{Dark: "#55ff55", Light: "#ff55ff"},
				BranchGrouping:      Colour{Dark: "#bbbb00", Light: "#0000bb"},
				BranchRemote:        Colour{Dark: "#ff5555", Light: "#55ffff"},
				Colon:               Colour{Dark: "#bbbbbb"},
				AuthorAngledBracket: Colour{Dark: "#bbbbbb"},
				AuthorText:          Colour{Dark: "#bbbbbb"},
				AuthorValue:         Colour{Dark: "#bbbbbb"},
				DateText:            Colour{Dark: "#bbbbbb"},
				DateValue:           Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)
			info := th.Info()

			assert.Equal(t, tt.info.HashText, toColour(info.HashText), "HashText")
			assert.Equal(t, tt.info.HashValue, toColour(info.HashValue), "HashValue")
			assert.Equal(t, tt.info.BranchHead, toColour(info.BranchHead), "BranchHead")
			assert.Equal(t, tt.info.BranchLocal, toColour(info.BranchLocal), "BranchLocal")
			assert.Equal(t, tt.info.BranchGrouping, toColour(info.BranchGrouping), "BranchGrouping")
			assert.Equal(t, tt.info.BranchRemote, toColour(info.BranchRemote), "BranchRemote")
			assert.Equal(t, tt.info.Colon, toColour(info.Colon), "Colon")
			assert.Equal(t, tt.info.AuthorAngledBracket, toColour(info.AuthorAngledBracket), "AuthorAngledBracket")
			assert.Equal(t, tt.info.AuthorText, toColour(info.AuthorText), "AuthorText")
			assert.Equal(t, tt.info.AuthorValue, toColour(info.AuthorValue), "AuthorValue")
			assert.Equal(t, tt.info.DateText, toColour(info.DateText), "DateText")
			assert.Equal(t, tt.info.DateValue, toColour(info.DateValue), "DateValue")
		})
	}
}

func TestMessage(t *testing.T) {
	tests := []struct {
		name    string
		message message
	}{
		{
			name: "Message",
			message: message{
				Summary: Colour{Dark: "#bbbbbb"},
				Body:    Colour{Dark: "#bbbbbb"},
				Footer:  Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)
			message := th.Message()

			assert.Equal(t, tt.message.Summary, toColour(message.Summary), "Summary")
			assert.Equal(t, tt.message.Body, toColour(message.Body), "Body")
			assert.Equal(t, tt.message.Footer, toColour(message.Footer), "Footer")
		})
	}
}

func TestShortcut(t *testing.T) {
	tests := []struct {
		name     string
		shortcut shortcut
	}{
		{
			name: "Shortcut",
			shortcut: shortcut{
				Key:          Colour{Dark: "#00bbbb", Light: "#bb0000"},
				Label:        Colour{Dark: "#00bb00", Light: "#bb00bb"},
				Plus:         Colour{Dark: "#bbbbbb"},
				AngleBracket: Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(config.ColourAdaptive)
			shortcut := th.Shortcut()

			assert.Equal(t, tt.shortcut.Key, toColour(shortcut.Key), "Key")
			assert.Equal(t, tt.shortcut.Label, toColour(shortcut.Label), "Label")
			assert.Equal(t, tt.shortcut.Plus, toColour(shortcut.Plus), "Plus")
			assert.Equal(t, tt.shortcut.AngleBracket, toColour(shortcut.AngleBracket), "AngleBracket")
		})
	}
}

func toColour(tc lipgloss.TerminalColor) Colour {
	switch clr := tc.(type) {
	case lipgloss.AdaptiveColor:
		return Colour{Dark: clr.Dark, Light: clr.Light}
	default:
		return Colour{Dark: fmt.Sprint(clr)}
	}
}
