package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type FooterModel struct {
	config FooterConfig
	focus  bool
}

type FooterConfig struct {
	name  string
	email string
}

func NewFooter(cfg commit.Config) FooterModel {
	c := FooterConfig{
		name:  cfg.Name,
		email: cfg.Email,
	}

	return FooterModel{
		config: c,
	}
}

func (m FooterModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m FooterModel) Update(msg tea.Msg) (FooterModel, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m FooterModel) View() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(m.footerRow())
}

func (m FooterModel) footerRow() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.signoff(),
	)
}

func (m FooterModel) signoff() string {
	s := colour("Signed-off-by", white)
	n := colour(m.config.name, white)
	e := colour(m.config.email, white)

	str := fmt.Sprintf("%s: %s <%s>", s, n, e)

	return lipgloss.NewStyle().
		Width(74).
		Height(1).
		MarginLeft(4).
		Align(lipgloss.Left, lipgloss.Center).
		Border(lipgloss.HiddenBorder(), false, true).
		Padding(0, 1, 0, 1).
		Render(str)
}
