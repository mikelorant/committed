package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
)

type HeaderModel struct {
	hash         string
	localBranch  string
	remoteBranch string
	branchRefs   []string
	remotes      []string
	name         string
	email        string
}

const (
	dateTimeFormat string = "Mon Jan 2 15:04:05 2006 -0700"
)

func NewHeader(cfg commit.Config) HeaderModel {
	return HeaderModel{
		hash:         cfg.Hash,
		localBranch:  cfg.LocalBranch,
		remoteBranch: cfg.RemoteBranch,
		branchRefs:   cfg.BranchRefs,
		remotes:      cfg.Remotes,
		name:         cfg.Name,
		email:        cfg.Email,
	}
}

func (m HeaderModel) render() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(headerColumn(
			m.hash,
			m.localBranch,
			m.remoteBranch,
			m.branchRefs,
			m.remotes,
			m.name,
			m.email,
		))
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

func hash(str string) string {
	k := colour("commit", yellow)
	h := colour(str, yellow)

	return lipgloss.NewStyle().
		MarginRight(1).
		Render(fmt.Sprintf("%s %s", k, h))
}

func branchRefs(lb, rb string, brefs, remotes []string) string {
	h := colour("HEAD ->", brightCyan, WithBold(true))

	l := colour(lb, brightGreen, WithBold(true))

	lp := colour("(", yellow)
	rp := colour(")", yellow)
	c := colour(",", yellow)

	str := fmt.Sprintf("%s %s", h, l)

	if rb != "" {
		str += fmt.Sprintf("%s %s", c, colour(rb, red, WithBold(true)))
	}

	for _, ref := range brefs {
		if containsPrefixes(ref, remotes) {
			rc := colour(ref, red, WithBold(true))
			str += fmt.Sprintf("%s %s", c, rc)
			continue
		}

		rc := colour(ref, brightGreen, WithBold(true))
		str += fmt.Sprintf("%s %s", c, rc)
	}

	return fmt.Sprintf("%s%s%s", lp, str, rp)
}

func author(name, email string) string {
	k := colour("author", white)
	n := colour(name, white)
	e := colour(email, white)

	return fmt.Sprintf("%s: %s <%s>", k, n, e)
}

func date(str string) string {
	k := colour("date", white)
	d := colour(str, white)

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
