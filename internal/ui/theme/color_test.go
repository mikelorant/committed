package theme_test

import (
	"fmt"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/ui/theme"

	"github.com/stretchr/testify/assert"
)

type body struct {
	Boundary            string
	TextAreaPlaceholder string
	TextAreaPrompt      string
	TextAreaFocusedText string
	TextAreaBlurredText string
	TextAreaCursorStyle string
}

type filterlist struct {
	Boundary                  string
	ListNormalTitle           string
	ListSelectedTitle         string
	ListNoItems               string
	TextInputPromptMark       string
	TextInputPromptText       string
	PaginatorDots             string
	TextInputPromptStyle      string
	TextInputTextStyle        string
	TextInputPlaceholderStyle string
	TextInputCursorStyle      string
}

type footer struct {
	View string
}

type header struct {
	EmojiBoundary                string
	SummaryBoundary              string
	CounterDivider               string
	CounterLimit                 string
	SummaryInputPromptStyle      string
	SummaryInputTextStyle        string
	SummaryInputPlaceholderStyle string
	SummaryInputCursorStyle      string
	CounterDefault               string
	CounterLow                   string
	CounterNormal                string
	CounterWarning               string
	CounterHigh                  string
}

type help struct {
	Boundary string
	Viewport string
}

type info struct {
	HashText            string
	HashValue           string
	BranchHead          string
	BranchLocal         string
	BranchGrouping      string
	BranchRemote        string
	Colon               string
	AuthorAngledBracket string
	AuthorText          string
	AuthorValue         string
	DateText            string
	DateValue           string
}

type message struct {
	Summary string
	Body    string
	Footer  string
}

type shortcut struct {
	Key          string
	Label        string
	Plus         string
	AngleBracket string
}

func TestBody(t *testing.T) {
	tests := []struct {
		name string
		body body
	}{
		{
			name: "body",
			body: body{
				Boundary:            "#bbbbbb",
				TextAreaPlaceholder: "#555555",
				TextAreaPrompt:      "#bbbbbb",
				TextAreaFocusedText: "#bbbbbb",
				TextAreaBlurredText: "#bbbbbb",
				TextAreaCursorStyle: "#bbbbbb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := theme.Body()

			assert.Equal(t, tt.body.Boundary, toHex(body.Boundary), "Boundary")
			assert.Equal(t, tt.body.TextAreaPlaceholder, toHex(body.TextAreaPlaceholder), "TextAreaPlaceholder")
			assert.Equal(t, tt.body.TextAreaPrompt, toHex(body.TextAreaPrompt), "TextAreaPrompt")
			assert.Equal(t, tt.body.TextAreaFocusedText, toHex(body.TextAreaFocusedText), "TextAreaFocusedText")
			assert.Equal(t, tt.body.TextAreaBlurredText, toHex(body.TextAreaBlurredText), "TextAreaBlurredText")
			assert.Equal(t, tt.body.TextAreaCursorStyle, toHex(body.TextAreaCursorStyle), "TextAreaCursorStyle")
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
				Boundary:                  "#bbbbbb",
				ListNormalTitle:           "#bbbbbb",
				ListSelectedTitle:         "#00bbbb",
				ListNoItems:               "#555555",
				TextInputPromptMark:       "#00bb00",
				TextInputPromptText:       "#bbbbbb",
				PaginatorDots:             "#00bbbb",
				TextInputPromptStyle:      "#bbbbbb",
				TextInputTextStyle:        "#bbbbbb",
				TextInputPlaceholderStyle: "#555555",
				TextInputCursorStyle:      "#bbbbbb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterlist := theme.FilterList()

			assert.Equal(t, tt.filterlist.Boundary, toHex(filterlist.Boundary), "Boundary")
			assert.Equal(t, tt.filterlist.ListNormalTitle, toHex(filterlist.ListNormalTitle), "ListNormalTitle")
			assert.Equal(t, tt.filterlist.ListSelectedTitle, toHex(filterlist.ListSelectedTitle), "ListSelectedTitle")
			assert.Equal(t, tt.filterlist.ListNoItems, toHex(filterlist.ListNoItems), "ListNoItems")
			assert.Equal(t, tt.filterlist.TextInputPromptMark, toHex(filterlist.TextInputPromptMark), "TextInputPromptMark")
			assert.Equal(t, tt.filterlist.TextInputPromptText, toHex(filterlist.TextInputPromptText), "TextInputPromptText")
			assert.Equal(t, tt.filterlist.PaginatorDots, toHex(filterlist.PaginatorDots), "PaginatorDots")
			assert.Equal(t, tt.filterlist.TextInputPromptStyle, toHex(filterlist.TextInputPromptStyle), "TextInputPromptStyle")
			assert.Equal(t, tt.filterlist.TextInputTextStyle, toHex(filterlist.TextInputTextStyle), "TextInputTextStyle")
			assert.Equal(t, tt.filterlist.TextInputPlaceholderStyle, toHex(filterlist.TextInputPlaceholderStyle), "TextInputPlaceholderStyle")
			assert.Equal(t, tt.filterlist.TextInputCursorStyle, toHex(filterlist.TextInputCursorStyle), "TextInputCursorStyle")
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
				View: "#bbbbbb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			footer := theme.Footer()

			assert.Equal(t, tt.footer.View, toHex(footer.View), "Boundary")
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
				EmojiBoundary:                "#bbbbbb",
				SummaryBoundary:              "#bbbbbb",
				CounterDivider:               "#bbbbbb",
				CounterLimit:                 "#bbbbbb",
				SummaryInputPromptStyle:      "#bbbbbb",
				SummaryInputTextStyle:        "#bbbbbb",
				SummaryInputPlaceholderStyle: "#555555",
				SummaryInputCursorStyle:      "#bbbbbb",
				CounterDefault:               "#bbbbbb",
				CounterLow:                   "#bbbb00",
				CounterNormal:                "#00bb00",
				CounterWarning:               "#bbbb00",
				CounterHigh:                  "#ff5555",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := theme.Header()

			assert.Equal(t, tt.header.EmojiBoundary, toHex(header.EmojiBoundary), "EmojiBoundary")
			assert.Equal(t, tt.header.SummaryBoundary, toHex(header.SummaryBoundary), "SummaryBoundary")
			assert.Equal(t, tt.header.CounterDivider, toHex(header.CounterDivider), "CounterDivider")
			assert.Equal(t, tt.header.CounterLimit, toHex(header.CounterLimit), "CounterLimit")
			assert.Equal(t, tt.header.SummaryInputPromptStyle, toHex(header.SummaryInputPromptStyle), "SummaryInputPromptStyle")
			assert.Equal(t, tt.header.SummaryInputTextStyle, toHex(header.SummaryInputTextStyle), "SummaryInputTextStyle")
			assert.Equal(t, tt.header.SummaryInputPlaceholderStyle, toHex(header.SummaryInputPlaceholderStyle), "SummaryInputPlaceholderStyle")
			assert.Equal(t, tt.header.SummaryInputCursorStyle, toHex(header.SummaryInputCursorStyle), "SummaryInputCursorStyle")
			assert.Equal(t, tt.header.CounterDefault, toHex(header.CounterDefault), "CounterDefault")
			assert.Equal(t, tt.header.CounterLow, toHex(header.CounterLow), "CounterLow")
			assert.Equal(t, tt.header.CounterNormal, toHex(header.CounterNormal), "CounterNormal")
			assert.Equal(t, tt.header.CounterWarning, toHex(header.CounterWarning), "CounterWarning")
			assert.Equal(t, tt.header.CounterHigh, toHex(header.CounterHigh), "CounterHigh")
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
				Boundary: "#bbbbbb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			help := theme.Help()

			assert.Equal(t, tt.help.Boundary, toHex(help.Boundary), "Boundary")
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
				HashText:            "#bbbb00",
				HashValue:           "#bbbb00",
				BranchHead:          "#55ffff",
				BranchLocal:         "#55ff55",
				BranchGrouping:      "#bbbb00",
				BranchRemote:        "#ff5555",
				Colon:               "#bbbbbb",
				AuthorAngledBracket: "#bbbbbb",
				AuthorText:          "#bbbbbb",
				AuthorValue:         "#bbbbbb",
				DateText:            "#bbbbbb",
				DateValue:           "#bbbbbb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := theme.Info()

			assert.Equal(t, tt.info.HashText, toHex(info.HashText), "HashText")
			assert.Equal(t, tt.info.HashValue, toHex(info.HashValue), "HashValue")
			assert.Equal(t, tt.info.BranchHead, toHex(info.BranchHead), "BranchHead")
			assert.Equal(t, tt.info.BranchLocal, toHex(info.BranchLocal), "BranchLocal")
			assert.Equal(t, tt.info.BranchGrouping, toHex(info.BranchGrouping), "BranchGrouping")
			assert.Equal(t, tt.info.BranchRemote, toHex(info.BranchRemote), "BranchRemote")
			assert.Equal(t, tt.info.Colon, toHex(info.Colon), "Colon")
			assert.Equal(t, tt.info.AuthorAngledBracket, toHex(info.AuthorAngledBracket), "AuthorAngledBracket")
			assert.Equal(t, tt.info.AuthorText, toHex(info.AuthorText), "AuthorText")
			assert.Equal(t, tt.info.AuthorValue, toHex(info.AuthorValue), "AuthorValue")
			assert.Equal(t, tt.info.DateText, toHex(info.DateText), "DateText")
			assert.Equal(t, tt.info.DateValue, toHex(info.DateValue), "DateValue")
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
				Summary: "#bbbbbb",
				Body:    "#bbbbbb",
				Footer:  "#bbbbbb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := theme.Message()

			assert.Equal(t, tt.message.Summary, toHex(message.Summary), "Summary")
			assert.Equal(t, tt.message.Body, toHex(message.Body), "Body")
			assert.Equal(t, tt.message.Footer, toHex(message.Footer), "Footer")
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
				Key:          "#00bbbb",
				Label:        "#00bb00",
				Plus:         "#bbbbbb",
				AngleBracket: "#bbbbbb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortcut := theme.Shortcut()

			assert.Equal(t, tt.shortcut.Key, toHex(shortcut.Key), "Key")
			assert.Equal(t, tt.shortcut.Label, toHex(shortcut.Label), "Label")
			assert.Equal(t, tt.shortcut.Plus, toHex(shortcut.Plus), "Plus")
			assert.Equal(t, tt.shortcut.AngleBracket, toHex(shortcut.AngleBracket), "AngleBracket")
		})
	}
}

func toHex(tc lipgloss.TerminalColor) string {
	return fmt.Sprintf("%v", tc)
}
