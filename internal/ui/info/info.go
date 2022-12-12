package info

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type Model struct {
	Hash         string
	LocalBranch  string
	RemoteBranch string
	BranchRefs   []string
	Remotes      []string
	Name         string
	Email        string
	Date         string
}

const (
	dateTimeFormat string = "Mon Jan 2 15:04:05 2006 -0700"

	black         = "0"
	red           = "1"
	green         = "2"
	yellow        = "3"
	blue          = "4"
	magenta       = "5"
	cyan          = "6"
	white         = "7"
	brightBlack   = "8"
	brightRed     = "9"
	brightGreen   = "10"
	brightYellow  = "11"
	brightBlue    = "12"
	brightMagenta = "13"
	brightCyan    = "14"
	brightWhite   = "15"
)

func New(cfg commit.Config) Model {
	return Model{
		Hash:         cfg.Hash,
		LocalBranch:  cfg.LocalBranch,
		RemoteBranch: cfg.RemoteBranch,
		BranchRefs:   cfg.BranchRefs,
		Remotes:      cfg.Remotes,
		Name:         cfg.Name,
		Email:        cfg.Email,
		Date:         time.Now().Format(dateTimeFormat),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(m.infoColumn())
}

func (m Model) infoColumn() string {
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

func (m Model) hash() string {
	k := lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow)).
		Render("commit")

	h := lipgloss.NewStyle().
		Foreground(lipgloss.Color(yellow)).
		Render(m.Hash)

	return lipgloss.NewStyle().
		MarginRight(1).
		Render(fmt.Sprintf("%s %s", k, h))
}

func (m Model) branchRefs() string {
	h := lipgloss.NewStyle().
		Foreground(lipgloss.Color(brightCyan)).
		Bold(true).
		Render("HEAD ->")

	l := lipgloss.NewStyle().
		Foreground(lipgloss.Color(brightGreen)).
		Bold(true).
		Render(m.LocalBranch)

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

	if m.RemoteBranch != "" {
		b := lipgloss.NewStyle().
			Foreground(lipgloss.Color(red)).
			Bold(true).
			Render(m.RemoteBranch)

		str += fmt.Sprintf("%s %s", c, b)
	}

	for _, ref := range m.BranchRefs {
		if containsPrefixes(ref, m.Remotes) {
			rc := lipgloss.NewStyle().
				Foreground(lipgloss.Color(red)).
				Bold(true).
				Render(m.RemoteBranch)
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

func (m Model) author() string {
	k := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render("author")
	n := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render(m.Name)
	e := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render(m.Email)

	return fmt.Sprintf("%s: %s <%s>", k, n, e)
}

func (m Model) date() string {
	k := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render("date")
	d := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Render(m.Date)

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
