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
	ReadyError                   lipgloss.TerminalColor
	ReadyIncomplete              lipgloss.TerminalColor
	ReadyOK                      lipgloss.TerminalColor
	CommitTypeNew                lipgloss.TerminalColor
	CommitTypeAmend              lipgloss.TerminalColor
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
func (t *Theme) Body() body {
	return body{
		Boundary:            t.Registry.Fg(),
		TextAreaPlaceholder: ToAdaptive(t.Registry.BrightBlack()),
		TextAreaPrompt:      t.Registry.Fg(),
		TextAreaFocusedText: t.Registry.Fg(),
		TextAreaBlurredText: t.Registry.Fg(),
		TextAreaCursorStyle: t.Registry.Fg(),
	}
}

//nolint:revive
func (t *Theme) FilterList() filterlist {
	return filterlist{
		Boundary:                  t.Registry.Fg(),
		ListNormalTitle:           t.Registry.Fg(),
		ListSelectedTitle:         ToAdaptive(t.Registry.Cyan()),
		ListNoItems:               ToAdaptive(t.Registry.BrightBlack()),
		TextInputPromptMark:       ToAdaptive(t.Registry.Green()),
		TextInputPromptText:       t.Registry.Fg(),
		PaginatorDots:             ToAdaptive(t.Registry.Cyan()),
		TextInputPromptStyle:      t.Registry.Fg(),
		TextInputTextStyle:        t.Registry.Fg(),
		TextInputPlaceholderStyle: ToAdaptive(t.Registry.BrightBlack()),
		TextInputCursorStyle:      t.Registry.Fg(),
	}
}

//nolint:revive
func (t *Theme) Footer() footer {
	return footer{
		View: t.Registry.Fg(),
	}
}

//nolint:revive
func (t *Theme) Header() header {
	return header{
		EmojiBoundary:                t.Registry.Fg(),
		SummaryBoundary:              t.Registry.Fg(),
		CounterDivider:               t.Registry.Fg(),
		CounterLimit:                 t.Registry.Fg(),
		SummaryInputPromptStyle:      t.Registry.Fg(),
		SummaryInputTextStyle:        t.Registry.Fg(),
		SummaryInputPlaceholderStyle: ToAdaptive(t.Registry.BrightBlack()),
		SummaryInputCursorStyle:      t.Registry.Fg(),
		CounterDefault:               t.Registry.Fg(),
		CounterLow:                   ToAdaptive(t.Registry.Yellow()),
		CounterNormal:                ToAdaptive(t.Registry.Green()),
		CounterWarning:               ToAdaptive(t.Registry.Yellow()),
		CounterHigh:                  ToAdaptive(t.Registry.BrightRed()),
		ReadyError:                   ToAdaptive(t.Registry.BrightRed()),
		ReadyIncomplete:              ToAdaptive(t.Registry.Yellow()),
		ReadyOK:                      ToAdaptive(t.Registry.Green()),
		CommitTypeNew:                ToAdaptive(t.Registry.Green()),
		CommitTypeAmend:              ToAdaptive(t.Registry.Yellow()),
	}
}

//nolint:revive
func (t *Theme) Help() help {
	return help{
		Boundary: t.Registry.Fg(),
		Viewport: t.Registry.Fg(),
	}
}

//nolint:revive
func (t *Theme) Info() info {
	return info{
		HashText:            ToAdaptive(t.Registry.Yellow()),
		HashValue:           ToAdaptive(t.Registry.Yellow()),
		BranchHead:          ToAdaptive(t.Registry.BrightCyan()),
		BranchLocal:         ToAdaptive(t.Registry.BrightGreen()),
		BranchGrouping:      ToAdaptive(t.Registry.Yellow()),
		BranchRemote:        ToAdaptive(t.Registry.BrightRed()),
		Colon:               t.Registry.Fg(),
		AuthorAngledBracket: t.Registry.Fg(),
		AuthorText:          t.Registry.Fg(),
		AuthorValue:         t.Registry.Fg(),
		DateText:            t.Registry.Fg(),
		DateValue:           t.Registry.Fg(),
	}
}

//nolint:revive
func (t *Theme) Message() message {
	return message{
		Summary: t.Registry.Fg(),
		Body:    t.Registry.Fg(),
		Footer:  t.Registry.Fg(),
	}
}

//nolint:revive
func (t *Theme) Shortcut() shortcut {
	return shortcut{
		Key:          ToAdaptive(t.Registry.Cyan()),
		Label:        ToAdaptive(t.Registry.Green()),
		Plus:         t.Registry.Fg(),
		AngleBracket: t.Registry.Fg(),
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
