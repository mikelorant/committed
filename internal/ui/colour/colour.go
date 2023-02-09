package colour

import (
	"fmt"
	"image/color"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
	"github.com/mikelorant/committed/internal/theme"
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

type Colour struct {
	registry *tint.Registry
}

type Msg int

func New(th theme.Theme) *Colour {
	return &Colour{
		registry: th.Registry,
	}
}

//nolint:revive
func (c *Colour) Body() body {
	clr := c.registry

	return body{
		Boundary:            clr.Fg(),
		TextAreaPlaceholder: ToAdaptive(clr.BrightBlack()),
		TextAreaPrompt:      clr.Fg(),
		TextAreaFocusedText: clr.Fg(),
		TextAreaBlurredText: clr.Fg(),
		TextAreaCursorStyle: clr.Fg(),
	}
}

//nolint:revive
func (c *Colour) FilterList() filterlist {
	clr := c.registry

	return filterlist{
		Boundary:                  clr.Fg(),
		ListNormalTitle:           clr.Fg(),
		ListSelectedTitle:         ToAdaptive(clr.Cyan()),
		ListNoItems:               ToAdaptive(clr.BrightBlack()),
		TextInputPromptMark:       ToAdaptive(clr.Green()),
		TextInputPromptText:       clr.Fg(),
		PaginatorDots:             ToAdaptive(clr.Cyan()),
		TextInputPromptStyle:      clr.Fg(),
		TextInputTextStyle:        clr.Fg(),
		TextInputPlaceholderStyle: ToAdaptive(clr.BrightBlack()),
		TextInputCursorStyle:      clr.Fg(),
	}
}

//nolint:revive
func (c *Colour) Footer() footer {
	clr := c.registry

	return footer{
		View: clr.Fg(),
	}
}

//nolint:revive
func (c *Colour) Header() header {
	clr := c.registry

	return header{
		EmojiBoundary:                clr.Fg(),
		SummaryBoundary:              clr.Fg(),
		CounterDivider:               clr.Fg(),
		CounterLimit:                 clr.Fg(),
		SummaryInputPromptStyle:      clr.Fg(),
		SummaryInputTextStyle:        clr.Fg(),
		SummaryInputPlaceholderStyle: ToAdaptive(clr.BrightBlack()),
		SummaryInputCursorStyle:      clr.Fg(),
		CounterDefault:               clr.Fg(),
		CounterLow:                   ToAdaptive(clr.Yellow()),
		CounterNormal:                ToAdaptive(clr.Green()),
		CounterWarning:               ToAdaptive(clr.Yellow()),
		CounterHigh:                  ToAdaptive(clr.BrightRed()),
		ReadyError:                   ToAdaptive(clr.BrightRed()),
		ReadyIncomplete:              ToAdaptive(clr.Yellow()),
		ReadyOK:                      ToAdaptive(clr.Green()),
		CommitTypeNew:                ToAdaptive(clr.Green()),
		CommitTypeAmend:              ToAdaptive(clr.Yellow()),
	}
}

//nolint:revive
func (c *Colour) Help() help {
	clr := c.registry

	return help{
		Boundary: clr.Fg(),
		Viewport: clr.Fg(),
	}
}

//nolint:revive
func (c *Colour) Info() info {
	clr := c.registry

	return info{
		HashText:            ToAdaptive(clr.Yellow()),
		HashValue:           ToAdaptive(clr.Yellow()),
		BranchHead:          ToAdaptive(clr.BrightCyan()),
		BranchLocal:         ToAdaptive(clr.BrightGreen()),
		BranchGrouping:      ToAdaptive(clr.Yellow()),
		BranchRemote:        ToAdaptive(clr.BrightRed()),
		Colon:               clr.Fg(),
		AuthorAngledBracket: clr.Fg(),
		AuthorText:          clr.Fg(),
		AuthorValue:         clr.Fg(),
		DateText:            clr.Fg(),
		DateValue:           clr.Fg(),
	}
}

//nolint:revive
func (c *Colour) Message() message {
	clr := c.registry

	return message{
		Summary: clr.Fg(),
		Body:    clr.Fg(),
		Footer:  clr.Fg(),
	}
}

//nolint:revive
func (c *Colour) Shortcut() shortcut {
	clr := c.registry

	return shortcut{
		Key:          ToAdaptive(clr.Cyan()),
		Label:        ToAdaptive(clr.Green()),
		Plus:         clr.Fg(),
		AngleBracket: clr.Fg(),
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

//nolint:ireturn
func Update() tea.Msg {
	var msg Msg

	return msg
}
