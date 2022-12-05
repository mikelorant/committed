package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type BodyModel struct {
	body string
}

func NewBody(cfg commit.Config) BodyModel {
	return BodyModel{
		body: cfg.Body,
	}
}

func (m BodyModel) render() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(body(m.body))
}

func body(str string) string {
	return lipgloss.NewStyle().
		Width(74).
		Height(19).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Top).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1, 0, 1).
		Faint(true).
		Render(strings.TrimSpace(str))
}
