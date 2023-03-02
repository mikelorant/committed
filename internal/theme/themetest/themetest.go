package themetest

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
)

type StubTint struct {
	id          string
	displayName string
}

func NewStubTint(id, displayName string) StubTint {
	return StubTint{
		id:          id,
		displayName: displayName,
	}
}

func NewStubTints(n int) []tint.Tint {
	tints := make([]tint.Tint, n)

	for i := 0; i < n; i++ {
		id := fmt.Sprintf("id%v", i)
		disp := fmt.Sprintf("%v", i)

		tints[i] = NewStubTint(id, disp)
	}

	return tints
}

func (t StubTint) DisplayName() string {
	return t.displayName
}

func (t StubTint) ID() string {
	return t.id
}

func (t StubTint) About() string {
	return ""
}

//nolint:ireturn
func (t StubTint) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Cursor() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Black() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Green() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Red() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) White() lipgloss.TerminalColor {
	return lipgloss.Color("")
}

//nolint:ireturn
func (t StubTint) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("")
}
