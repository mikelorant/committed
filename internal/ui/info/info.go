package info

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/fuzzy"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/ui/filterlist"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	Expand        bool
	DefaultHeight int
	ExpandHeight  int
	Hash          string
	LocalBranch   string
	RemoteBranch  string
	BranchRefs    []string
	Remotes       []string
	Date          string
	Author        repository.User
	Authors       []repository.User

	focus      bool
	styles     Styles
	filterList filterlist.Model
}

const (
	filterPromptText = "Choose an author:"
	filterHeight     = 3
)

func New(state *commit.State) Model {
	if len(state.Repository.Users) == 0 {
		state.Repository.Users = []repository.User{{}}
	}

	m := Model{
		Hash:         state.Placeholders.Hash,
		LocalBranch:  state.Repository.Branch.Local,
		RemoteBranch: state.Repository.Branch.Remote,
		BranchRefs:   state.Repository.Branch.Refs,
		Remotes:      state.Repository.Remotes,
		Date:         time.Now().Format(dateTimeFormat),
		Author:       state.Repository.Users[0],
		Authors:      state.Repository.Users,
		styles:       defaultStyles(),
		filterList: filterlist.New(
			castToListItems(state.Repository.Users),
			filterPromptText,
			filterHeight,
		),
	}

	if state.Amend {
		m.Hash = state.Repository.Head.Hash
		m.Date = state.Repository.Head.When.Format(dateTimeFormat)
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.focus {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.Author = m.filterList.SelectedItem().(listItem).author
			}
		}
	}

	//nolint:gocritic
	switch msg.(type) {
	case theme.Msg:
		m.styles = defaultStyles()
	}

	switch {
	case !m.focus && m.filterList.Focused():
		m.filterList.Blur()
	case m.focus && !m.filterList.Focused():
		m.filterList.Focus()
		fallthrough
	case m.focus:
		ranks := fuzzy.Rank(m.filterList.Filter(), castToFuzzyItems(m.Authors))

		items := make([]list.Item, len(ranks))
		for i, rank := range ranks {
			items[i] = castToListItems(m.Authors)[rank]
		}
		m.filterList.SetItems(items)
	}

	m.filterList, cmd = filterlist.ToModel(m.filterList.Update(msg))
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.styles.infoBoundary.Render(m.infoColumn())
}

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m Model) Focused() bool {
	return m.focus
}

func (m Model) infoColumn() string {
	hashBranchRefs := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.hash(),
		m.branchRefs(),
	)

	it := lipgloss.JoinVertical(
		lipgloss.Top,
		hashBranchRefs,
		m.author(),
		m.date(),
	)

	if !m.Expand {
		return it
	}

	fl := m.styles.filterListBoundary.Render(m.filterList.View())

	return lipgloss.JoinVertical(
		lipgloss.Top,
		it,
		fl,
	)
}

func (m Model) hash() string {
	k := m.styles.hashText
	h := m.styles.hashValue.Render(m.Hash)

	return m.styles.hashBoundary.Render(fmt.Sprintf("%s %s", k, h))
}

func (m Model) branchRefs() string {
	if m.LocalBranch == "" {
		return ""
	}

	h := m.styles.branchHead
	l := m.styles.branchLocal.Render(m.LocalBranch)

	lp := m.styles.branchGrouping.Render("(")
	rp := m.styles.branchGrouping.Render(")")
	c := m.styles.branchGrouping.Render(",")

	str := fmt.Sprintf("%s %s", h, l)

	if m.RemoteBranch != "" {
		b := m.styles.branchRemote.Render(m.RemoteBranch)

		str += fmt.Sprintf("%s %s", c, b)
	}

	for _, ref := range m.BranchRefs {
		if containsPrefixes(ref, m.Remotes) {
			rc := m.styles.branchRemote.Render(m.RemoteBranch)
			str += fmt.Sprintf("%s %s", c, rc)
			continue
		}

		rc := m.styles.branchLocal.Render(ref)
		str += fmt.Sprintf("%s %s", c, rc)
	}

	return fmt.Sprintf("%s%s%s", lp, str, rp)
}

func (m Model) author() string {
	k := m.styles.authorText
	c := m.styles.colon
	lb := m.styles.authorAngledBracket.Render("<")
	rb := m.styles.authorAngledBracket.Render(">")
	n := m.styles.authorValue.Render(m.Author.Name)
	e := m.styles.authorValue.Render(m.Author.Email)

	return fmt.Sprintf("%s%s %s %s%s%s", k, c, n, lb, e, rb)
}

func (m Model) date() string {
	k := m.styles.dateText
	c := m.styles.colon
	d := m.styles.dateValue.Render(m.Date)

	return fmt.Sprintf("%s%s   %s", k, c, d)
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}

func containsPrefixes(str string, ps []string) bool {
	for _, p := range ps {
		if strings.HasPrefix(str, fmt.Sprintf("%s/", p)) {
			return true
		}
	}

	return false
}
