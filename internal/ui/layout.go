package ui

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

const (
	dateTimeFormat string = "Mon Jan 2 15:04:05 2006 -0700"
	subjectLimit   int    = 50
)

func commit(m model) string {
	blockHeader := lipgloss.NewStyle().
		MarginBottom(1).
		Render(headerColumn(m.commit, m.name, m.email))

	blockSubject := lipgloss.NewStyle().
		MarginBottom(1).
		Render(subjectRow(m.emoji, m.summary))

	blockBody := lipgloss.NewStyle().
		MarginBottom(1).
		Render(body(m.body))

	blockFooter := lipgloss.NewStyle().
		MarginBottom(1).
		Render(footerRow(m.name, m.email))

	blockStatus := lipgloss.NewStyle().
		MarginBottom(1).
		Render(statusRow())

	return lipgloss.JoinVertical(lipgloss.Top,
		blockHeader,
		blockSubject,
		blockBody,
		blockFooter,
		blockStatus,
	)
}

func headerColumn(h, n, e string) string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		hash(h),
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
