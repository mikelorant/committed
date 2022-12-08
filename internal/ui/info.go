package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type InfoModel struct {
	config InfoConfig
	focus  bool
}

type InfoConfig struct {
	hash         string
	localBranch  string
	remoteBranch string
	branchRefs   []string
	remotes      []string
	name         string
	email        string
	date         string
}

const (
	dateTimeFormat string = "Mon Jan 2 15:04:05 2006 -0700"
)

func NewInfo(cfg commit.Config) InfoModel {
	c := InfoConfig{
		hash:         cfg.Hash,
		localBranch:  cfg.LocalBranch,
		remoteBranch: cfg.RemoteBranch,
		branchRefs:   cfg.BranchRefs,
		remotes:      cfg.Remotes,
		name:         cfg.Name,
		email:        cfg.Email,
		date:         time.Now().Format(dateTimeFormat),
	}

	return InfoModel{
		config: c,
	}
}

func (m InfoModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m InfoModel) Update(msg tea.Msg) (InfoModel, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m InfoModel) View() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(m.infoColumn())
}

func (m InfoModel) infoColumn() string {
	hashBranchRefs := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.hash(),
		m.branchRefs(),
	)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		hashBranchRefs,
		m.author(),
		m.date(),
	)
}

func (m InfoModel) hash() string {
	k := lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow)).
		Render("commit")

	h := lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow)).
		Render(m.config.hash)

	return lipgloss.NewStyle().
		MarginRight(1).
		Render(fmt.Sprintf("%s %s", k, h))
}

func (m InfoModel) branchRefs() string {
	h := lipgloss.NewStyle().
		Foreground(lipgloss.Color(brightCyan)).
		Bold(true).
		Render("HEAD ->")

	l := lipgloss.NewStyle().
		Foreground(lipgloss.Color(brightGreen)).
		Bold(true).
		Render(m.config.localBranch)

	lp := lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow)).
		Render("(")

	rp := lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow)).
		Render(")")

	c := lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow)).
		Render(",")

	str := fmt.Sprintf("%s %s", h, l)

	if m.config.remoteBranch != "" {
		b := lipgloss.NewStyle().
			Foreground(lipgloss.Color(red)).
			Bold(true).
			Render(m.config.remoteBranch)

		str += fmt.Sprintf("%s %s", c, b)
	}

	for _, ref := range m.config.branchRefs {
		if containsPrefixes(ref, m.config.remotes) {
			rc := lipgloss.NewStyle().
				Foreground(lipgloss.Color(red)).
				Bold(true).
				Render(m.config.remoteBranch)
			str += fmt.Sprintf("%s %s", c, rc)
			continue
		}

		rc := lipgloss.NewStyle().
			Foreground(lipgloss.Color(brightGreen)).
			Bold(true).
			Render(ref)
		str += fmt.Sprintf("%s %s", c, rc)
	}

	return fmt.Sprintf("%s%s%s", lp, str, rp)
}

func (m InfoModel) author() string {
	k := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render("author")
	n := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render(m.config.name)
	e := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render(m.config.email)

	return fmt.Sprintf("%s: %s <%s>", k, n, e)
}

func (m InfoModel) date() string {
	k := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render("date")
	d := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render(m.config.date)

	return fmt.Sprintf("%s:   %s", k, d)
}

func containsPrefixes(str string, ps []string) bool {
	for _, p := range ps {
		if strings.HasPrefix(str, fmt.Sprintf("%s/", p)) {
			return true
		}
	}

	return false
}
