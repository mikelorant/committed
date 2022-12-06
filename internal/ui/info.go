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
	k := colour("commit", yellow)
	h := colour(m.config.hash, yellow)

	return lipgloss.NewStyle().
		MarginRight(1).
		Render(fmt.Sprintf("%s %s", k, h))
}

func (m InfoModel) branchRefs() string {
	h := colour("HEAD ->", brightCyan, WithBold(true))

	l := colour(m.config.localBranch, brightGreen, WithBold(true))

	lp := colour("(", yellow)
	rp := colour(")", yellow)
	c := colour(",", yellow)

	str := fmt.Sprintf("%s %s", h, l)

	if m.config.remoteBranch != "" {
		str += fmt.Sprintf("%s %s", c, colour(m.config.remoteBranch, red, WithBold(true)))
	}

	for _, ref := range m.config.branchRefs {
		if containsPrefixes(ref, m.config.remotes) {
			rc := colour(ref, red, WithBold(true))
			str += fmt.Sprintf("%s %s", c, rc)
			continue
		}

		rc := colour(ref, brightGreen, WithBold(true))
		str += fmt.Sprintf("%s %s", c, rc)
	}

	return fmt.Sprintf("%s%s%s", lp, str, rp)
}

func (m InfoModel) author() string {
	k := colour("author", white)
	n := colour(m.config.name, white)
	e := colour(m.config.email, white)

	return fmt.Sprintf("%s: %s <%s>", k, n, e)
}

func (m InfoModel) date() string {
	k := colour("date", white)
	d := colour(m.config.date, white)

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
