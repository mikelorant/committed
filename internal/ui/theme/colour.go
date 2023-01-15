package theme

import (
	"fmt"
	"image/color"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/gamut"
)

type body struct {
	Boundary            lipgloss.TerminalColor
	TextAreaPlaceholder lipgloss.TerminalColor
	TextAreaPrompt      lipgloss.TerminalColor
	TextAreaFocusedText lipgloss.TerminalColor
	TextAreaBlurredText lipgloss.TerminalColor
	TextAreaCursorStyle lipgloss.TerminalColor
}

type filterlist struct {
	Boundary                  lipgloss.TerminalColor
	ListNormalTitle           lipgloss.TerminalColor
	ListSelectedTitle         lipgloss.TerminalColor
	ListNoItems               lipgloss.TerminalColor
	TextInputPromptMark       lipgloss.TerminalColor
	TextInputPromptText       lipgloss.TerminalColor
	PaginatorDots             lipgloss.TerminalColor
	TextInputPromptStyle      lipgloss.TerminalColor
	TextInputTextStyle        lipgloss.TerminalColor
	TextInputPlaceholderStyle lipgloss.TerminalColor
	TextInputCursorStyle      lipgloss.TerminalColor
}

type footer struct {
	View lipgloss.TerminalColor
}

type header struct {
	EmojiBoundary                lipgloss.TerminalColor
	SummaryBoundary              lipgloss.TerminalColor
	CounterDivider               lipgloss.TerminalColor
	CounterLimit                 lipgloss.TerminalColor
	SummaryInputPromptStyle      lipgloss.TerminalColor
	SummaryInputTextStyle        lipgloss.TerminalColor
	SummaryInputPlaceholderStyle lipgloss.TerminalColor
	SummaryInputCursorStyle      lipgloss.TerminalColor
	CounterDefault               lipgloss.TerminalColor
	CounterLow                   lipgloss.TerminalColor
	CounterNormal                lipgloss.TerminalColor
	CounterWarning               lipgloss.TerminalColor
	CounterHigh                  lipgloss.TerminalColor
}

type help struct {
	Boundary lipgloss.TerminalColor
	Viewport lipgloss.TerminalColor
}

type info struct {
	HashText            lipgloss.TerminalColor
	HashValue           lipgloss.TerminalColor
	BranchHead          lipgloss.TerminalColor
	BranchLocal         lipgloss.TerminalColor
	BranchGrouping      lipgloss.TerminalColor
	BranchRemote        lipgloss.TerminalColor
	Colon               lipgloss.TerminalColor
	AuthorAngledBracket lipgloss.TerminalColor
	AuthorText          lipgloss.TerminalColor
	AuthorValue         lipgloss.TerminalColor
	DateText            lipgloss.TerminalColor
	DateValue           lipgloss.TerminalColor
}

type message struct {
	Summary lipgloss.TerminalColor
	Body    lipgloss.TerminalColor
	Footer  lipgloss.TerminalColor
}

type shortcut struct {
	Key          lipgloss.TerminalColor
	Label        lipgloss.TerminalColor
	Plus         lipgloss.TerminalColor
	AngleBracket lipgloss.TerminalColor
}

//nolint:revive
func Body() body {
	tint := Tint()

	return body{
		Boundary:            tint.Fg(),
		TextAreaPlaceholder: ToAdaptive(tint.BrightBlack()),
		TextAreaPrompt:      tint.Fg(),
		TextAreaFocusedText: tint.Fg(),
		TextAreaBlurredText: tint.Fg(),
		TextAreaCursorStyle: tint.Fg(),
	}
}

//nolint:revive
func FilterList() filterlist {
	tint := Tint()

	return filterlist{
		Boundary:                  tint.Fg(),
		ListNormalTitle:           tint.Fg(),
		ListSelectedTitle:         ToAdaptive(tint.Cyan()),
		ListNoItems:               ToAdaptive(tint.BrightBlack()),
		TextInputPromptMark:       ToAdaptive(tint.Green()),
		TextInputPromptText:       tint.Fg(),
		PaginatorDots:             ToAdaptive(tint.Cyan()),
		TextInputPromptStyle:      tint.Fg(),
		TextInputTextStyle:        tint.Fg(),
		TextInputPlaceholderStyle: ToAdaptive(tint.BrightBlack()),
		TextInputCursorStyle:      tint.Fg(),
	}
}

//nolint:revive
func Footer() footer {
	tint := Tint()

	return footer{
		View: tint.Fg(),
	}
}

//nolint:revive
func Header() header {
	tint := Tint()

	return header{
		EmojiBoundary:                tint.Fg(),
		SummaryBoundary:              tint.Fg(),
		CounterDivider:               tint.Fg(),
		CounterLimit:                 tint.Fg(),
		SummaryInputPromptStyle:      tint.Fg(),
		SummaryInputTextStyle:        tint.Fg(),
		SummaryInputPlaceholderStyle: ToAdaptive(tint.BrightBlack()),
		SummaryInputCursorStyle:      tint.Fg(),
		CounterDefault:               tint.Fg(),
		CounterLow:                   ToAdaptive(tint.Yellow()),
		CounterNormal:                ToAdaptive(tint.Green()),
		CounterWarning:               ToAdaptive(tint.Yellow()),
		CounterHigh:                  ToAdaptive(tint.BrightRed()),
	}
}

//nolint:revive
func Help() help {
	tint := Tint()

	return help{
		Boundary: tint.Fg(),
		Viewport: tint.Fg(),
	}
}

//nolint:revive
func Info() info {
	tint := Tint()

	return info{
		HashText:            ToAdaptive(tint.Yellow()),
		HashValue:           ToAdaptive(tint.Yellow()),
		BranchHead:          ToAdaptive(tint.BrightCyan()),
		BranchLocal:         ToAdaptive(tint.BrightGreen()),
		BranchGrouping:      ToAdaptive(tint.Yellow()),
		BranchRemote:        ToAdaptive(tint.BrightRed()),
		Colon:               tint.Fg(),
		AuthorAngledBracket: tint.Fg(),
		AuthorText:          tint.Fg(),
		AuthorValue:         tint.Fg(),
		DateText:            tint.Fg(),
		DateValue:           tint.Fg(),
	}
}

//nolint:revive
func Message() message {
	tint := Tint()

	return message{
		Summary: tint.Fg(),
		Body:    tint.Fg(),
		Footer:  tint.Fg(),
	}
}

//nolint:revive
func Shortcut() shortcut {
	tint := Tint()

	return shortcut{
		Key:          ToAdaptive(tint.Cyan()),
		Label:        ToAdaptive(tint.Green()),
		Plus:         tint.Fg(),
		AngleBracket: tint.Fg(),
	}
}

func ToAdaptive(clr color.Color) lipgloss.AdaptiveColor {
	return lipgloss.AdaptiveColor{
		Dark:  ToDefault(clr),
		Light: ToComplementary(ToDefault(clr)),
	}
}

func ToDefault(clr color.Color) string {
	return fmt.Sprint(clr)
}

func ToComplementary(hexClr string) string {
	clr := gamut.Hex(hexClr)
	compClr := gamut.Complementary(clr)

	return gamut.ToHex(compClr)
}
