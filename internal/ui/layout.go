package ui

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

const (
	dateTimeFormat string = "Mon Jan 2 15:04:05 2006 -0700"
	subjectLimit   int    = 50
)

func commitBlock(m model) string {
	headerBlock := lipgloss.NewStyle().
		MarginBottom(1).
		Render(headerColumn(
			m.config.Hash,
			m.config.LocalBranch,
			m.config.RemoteBranch,
			m.config.BranchRefs,
			m.config.Remotes,
			m.config.Name,
			m.config.Email,
		))

	subjectBlock := lipgloss.NewStyle().
		MarginBottom(1).
		Render(subjectRow(m.config.Emoji, m.config.Summary))

	bodyBlock := lipgloss.NewStyle().
		MarginBottom(1).
		Render(body(m.config.Body))

	footerBlock := lipgloss.NewStyle().
		MarginBottom(1).
		Render(footerRow(m.config.Name, m.config.Email))

	statusBlock := lipgloss.NewStyle().
		MarginBottom(1).
		Render(statusRow())

	return lipgloss.JoinVertical(lipgloss.Top,
		headerBlock,
		subjectBlock,
		bodyBlock,
		footerBlock,
		statusBlock,
	)
}

func headerColumn(h, lb, rb string, brefs, remotes []string, n, e string) string {
	hashBranchRefs := lipgloss.JoinHorizontal(
		lipgloss.Top,
		hash(h),
		branchRefs(lb, rb, brefs, remotes),
	)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		hashBranchRefs,
		author(n, e),
		date(time.Now().Format(dateTimeFormat)),
	)
}

func subjectRow(e, s string) string {
	i := len(s)
	if e != "" {
		i += 2
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		emoji(e),
		summary(s),
		counter(i, subjectLimit),
	)
}

func footerRow(n, e string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		signoff(n, e),
	)
}

func statusRow() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		modifier(),
		commands(shortcuts),
	)
}
