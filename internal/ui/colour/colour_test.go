package colour_test

import (
	"fmt"
	"testing"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

type Colour struct {
	Dark  string
	Light string
}

type body struct {
	Boundary            Colour
	FocusBoundary       Colour
	TextAreaPlaceholder Colour
	TextAreaPrompt      Colour
	TextAreaFocusedText Colour
	TextAreaBlurredText Colour
	TextAreaCursorStyle Colour
}

type filterlist struct {
	Boundary                  Colour
	FocusBoundary             Colour
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
	EmojiFocusBoundary           Colour
	SummaryBoundary              Colour
	SummaryFocusBoundary         Colour
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
	BranchTag           Colour
	Colon               Colour
	AuthorAngledBracket Colour
	AuthorText          Colour
	AuthorValue         Colour
	DateText            Colour
	DateValue           Colour
}

type message struct {
	Message Colour
}

type option struct {
	SectionBoundary      Colour
	SectionBoundaryFocus Colour
}

type optionSection struct {
	Category         Colour
	CategorySelected Colour
	CategorySpacer   Colour
	CategoryPrompt   Colour
	Setting          Colour
	SettingSelected  Colour
	SettingSpacer    Colour
	SettingPrompt    Colour
	SettingJoiner    Colour
}

type shortcut struct {
	Key          Colour
	Label        Colour
	Plus         Colour
	AngleBracket Colour
}

func TestBody(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		body body
	}{
		{
			name: "body",
			body: body{
				Boundary:            Colour{Dark: "#555555", Light: "#555555"},
				FocusBoundary:       Colour{Dark: "#bbbbbb"},
				TextAreaPlaceholder: Colour{Dark: "#555555", Light: "#555555"},
				TextAreaPrompt:      Colour{Dark: "#bbbbbb"},
				TextAreaFocusedText: Colour{Dark: "#bbbbbb"},
				TextAreaBlurredText: Colour{Dark: "#bbbbbb"},
				TextAreaCursorStyle: Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).Body()

			assert.Equal(t, tt.body.Boundary, toColour(clr.Boundary), "Boundary")
			assert.Equal(t, tt.body.FocusBoundary, toColour(clr.FocusBoundary), "FocusBoundary")
			assert.Equal(t, tt.body.TextAreaPlaceholder, toColour(clr.TextAreaPlaceholder), "TextAreaPlaceholder")
			assert.Equal(t, tt.body.TextAreaPrompt, toColour(clr.TextAreaPrompt), "TextAreaPrompt")
			assert.Equal(t, tt.body.TextAreaFocusedText, toColour(clr.TextAreaFocusedText), "TextAreaFocusedText")
			assert.Equal(t, tt.body.TextAreaBlurredText, toColour(clr.TextAreaBlurredText), "TextAreaBlurredText")
			assert.Equal(t, tt.body.TextAreaCursorStyle, toColour(clr.TextAreaCursorStyle), "TextAreaCursorStyle")
		})
	}
}

func TestFilterList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		filterlist filterlist
	}{
		{
			name: "filterlist",
			filterlist: filterlist{
				Boundary:                  Colour{Dark: "#555555", Light: "#555555"},
				FocusBoundary:             Colour{Dark: "#bbbbbb"},
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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).FilterList()

			assert.Equal(t, tt.filterlist.Boundary, toColour(clr.Boundary), "Boundary")
			assert.Equal(t, tt.filterlist.FocusBoundary, toColour(clr.FocusBoundary), "FocusBoundary")
			assert.Equal(t, tt.filterlist.ListNormalTitle, toColour(clr.ListNormalTitle), "ListNormalTitle")
			assert.Equal(t, tt.filterlist.ListSelectedTitle, toColour(clr.ListSelectedTitle), "ListSelectedTitle")
			assert.Equal(t, tt.filterlist.ListNoItems, toColour(clr.ListNoItems), "ListNoItems")
			assert.Equal(t, tt.filterlist.TextInputPromptMark, toColour(clr.TextInputPromptMark), "TextInputPromptMark")
			assert.Equal(t, tt.filterlist.TextInputPromptText, toColour(clr.TextInputPromptText), "TextInputPromptText")
			assert.Equal(t, tt.filterlist.PaginatorDots, toColour(clr.PaginatorDots), "PaginatorDots")
			assert.Equal(t, tt.filterlist.TextInputPromptStyle, toColour(clr.TextInputPromptStyle), "TextInputPromptStyle")
			assert.Equal(t, tt.filterlist.TextInputTextStyle, toColour(clr.TextInputTextStyle), "TextInputTextStyle")
			assert.Equal(t, tt.filterlist.TextInputPlaceholderStyle, toColour(clr.TextInputPlaceholderStyle), "TextInputPlaceholderStyle")
			assert.Equal(t, tt.filterlist.TextInputCursorStyle, toColour(clr.TextInputCursorStyle), "TextInputCursorStyle")
		})
	}
}

func TestFooter(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).Footer()

			assert.Equal(t, tt.footer.View, toColour(clr.View), "Boundary")
		})
	}
}

func TestHeader(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		header header
	}{
		{
			name: "Header",
			header: header{
				EmojiBoundary:                Colour{Dark: "#555555", Light: "#555555"},
				EmojiFocusBoundary:           Colour{Dark: "#bbbbbb"},
				SummaryBoundary:              Colour{Dark: "#555555", Light: "#555555"},
				SummaryFocusBoundary:         Colour{Dark: "#bbbbbb"},
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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).Header()

			assert.Equal(t, tt.header.EmojiBoundary, toColour(clr.EmojiBoundary), "EmojiBoundary")
			assert.Equal(t, tt.header.EmojiFocusBoundary, toColour(clr.EmojiFocusBoundary), "EmojiFocusBoundary")
			assert.Equal(t, tt.header.SummaryBoundary, toColour(clr.SummaryBoundary), "SummaryBoundary")
			assert.Equal(t, tt.header.SummaryFocusBoundary, toColour(clr.SummaryFocusBoundary), "SummaryFocusBoundary")
			assert.Equal(t, tt.header.CounterDivider, toColour(clr.CounterDivider), "CounterDivider")
			assert.Equal(t, tt.header.CounterLimit, toColour(clr.CounterLimit), "CounterLimit")
			assert.Equal(t, tt.header.SummaryInputPromptStyle, toColour(clr.SummaryInputPromptStyle), "SummaryInputPromptStyle")
			assert.Equal(t, tt.header.SummaryInputTextStyle, toColour(clr.SummaryInputTextStyle), "SummaryInputTextStyle")
			assert.Equal(t, tt.header.SummaryInputPlaceholderStyle, toColour(clr.SummaryInputPlaceholderStyle), "SummaryInputPlaceholderStyle")
			assert.Equal(t, tt.header.SummaryInputCursorStyle, toColour(clr.SummaryInputCursorStyle), "SummaryInputCursorStyle")
			assert.Equal(t, tt.header.CounterDefault, toColour(clr.CounterDefault), "CounterDefault")
			assert.Equal(t, tt.header.CounterLow, toColour(clr.CounterLow), "CounterLow")
			assert.Equal(t, tt.header.CounterNormal, toColour(clr.CounterNormal), "CounterNormal")
			assert.Equal(t, tt.header.CounterWarning, toColour(clr.CounterWarning), "CounterWarning")
			assert.Equal(t, tt.header.CounterHigh, toColour(clr.CounterHigh), "CounterHigh")
			assert.Equal(t, tt.header.CommitTypeNew, toColour(clr.CommitTypeNew), "CommitTypeNew")
			assert.Equal(t, tt.header.CommitTypeAmend, toColour(clr.CommitTypeAmend), "CommitTypeAmend")
		})
	}
}

func TestHelp(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).Help()

			assert.Equal(t, tt.help.Boundary, toColour(clr.Boundary), "Boundary")
		})
	}
}

func TestInfo(t *testing.T) {
	t.Parallel()

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
				BranchTag:           Colour{Dark: "#ffff55", Light: "#5555ff"},
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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).Info()

			assert.Equal(t, tt.info.HashText, toColour(clr.HashText), "HashText")
			assert.Equal(t, tt.info.HashValue, toColour(clr.HashValue), "HashValue")
			assert.Equal(t, tt.info.BranchHead, toColour(clr.BranchHead), "BranchHead")
			assert.Equal(t, tt.info.BranchLocal, toColour(clr.BranchLocal), "BranchLocal")
			assert.Equal(t, tt.info.BranchGrouping, toColour(clr.BranchGrouping), "BranchGrouping")
			assert.Equal(t, tt.info.BranchRemote, toColour(clr.BranchRemote), "BranchRemote")
			assert.Equal(t, tt.info.BranchTag, toColour(clr.BranchTag), "BranchTag")
			assert.Equal(t, tt.info.Colon, toColour(clr.Colon), "Colon")
			assert.Equal(t, tt.info.AuthorAngledBracket, toColour(clr.AuthorAngledBracket), "AuthorAngledBracket")
			assert.Equal(t, tt.info.AuthorText, toColour(clr.AuthorText), "AuthorText")
			assert.Equal(t, tt.info.AuthorValue, toColour(clr.AuthorValue), "AuthorValue")
			assert.Equal(t, tt.info.DateText, toColour(clr.DateText), "DateText")
			assert.Equal(t, tt.info.DateValue, toColour(clr.DateValue), "DateValue")
		})
	}
}

func TestMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message message
	}{
		{
			name: "Message",
			message: message{
				Message: Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).Message()

			assert.Equal(t, tt.message.Message, toColour(clr.Message), "Summary")
		})
	}
}

func TestOption(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		option option
	}{
		{
			name: "Option",
			option: option{
				SectionBoundary:      Colour{Dark: "#555555", Light: "#555555"},
				SectionBoundaryFocus: Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).Option()

			assert.Equal(t, tt.option.SectionBoundary, toColour(clr.SectionBoundary), "SectionBoundary")
			assert.Equal(t, tt.option.SectionBoundaryFocus, toColour(clr.SectionBoundaryFocus), "SectionBoundaryFocus")
		})
	}
}

func TestOptionSection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		optionSection optionSection
	}{
		{
			name: "OptionSection",
			optionSection: optionSection{
				Category:         Colour{Dark: "#bbbbbb"},
				CategorySelected: Colour{Dark: "#ffffff", Light: "#ffffff"},
				CategorySpacer:   Colour{Dark: "#bbbbbb"},
				CategoryPrompt:   Colour{Dark: "#00bbbb"},
				Setting:          Colour{Dark: "#bbbbbb"},
				SettingSelected:  Colour{Dark: "#ffffff", Light: "#ffffff"},
				SettingSpacer:    Colour{Dark: "#bbbbbb"},
				SettingPrompt:    Colour{Dark: "#bbbbbb"},
				SettingJoiner:    Colour{Dark: "#bbbbbb"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).OptionSection()

			assert.Equal(t, tt.optionSection.Category, toColour(clr.Category), "Category")
			assert.Equal(t, tt.optionSection.CategorySelected, toColour(clr.CategorySelected), "CategorySelected")
			assert.Equal(t, tt.optionSection.CategorySpacer, toColour(clr.CategorySpacer), "CategorySpacer")
			assert.Equal(t, tt.optionSection.CategoryPrompt, toColour(clr.CategoryPrompt), "CategoryPrompt")
			assert.Equal(t, tt.optionSection.Setting, toColour(clr.Setting), "Setting")
			assert.Equal(t, tt.optionSection.SettingSelected, toColour(clr.SettingSelected), "SettingSelected")
			assert.Equal(t, tt.optionSection.SettingSpacer, toColour(clr.SettingSpacer), "SettingSpacer")
			assert.Equal(t, tt.optionSection.SettingPrompt, toColour(clr.SettingPrompt), "SettingPrompt")
			assert.Equal(t, tt.optionSection.SettingJoiner, toColour(clr.SettingJoiner), "SettingJoiner")
		})
	}
}

func TestShortcut(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			clr := colour.New(theme.New(config.ColourAdaptive)).Shortcut()

			assert.Equal(t, tt.shortcut.Key, toColour(clr.Key), "Key")
			assert.Equal(t, tt.shortcut.Label, toColour(clr.Label), "Label")
			assert.Equal(t, tt.shortcut.Plus, toColour(clr.Plus), "Plus")
			assert.Equal(t, tt.shortcut.AngleBracket, toColour(clr.AngleBracket), "AngleBracket")
		})
	}
}

func toColour(clr lipgloss.TerminalColor) Colour {
	switch clr := clr.(type) {
	case lipgloss.AdaptiveColor:
		return Colour{Dark: clr.Dark, Light: clr.Light}
	default:
		return Colour{Dark: fmt.Sprint(clr)}
	}
}
