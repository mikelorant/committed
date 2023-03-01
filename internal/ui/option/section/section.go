package section

import (
	"fmt"
	"strings"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/colour"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/slices"
)

type Model struct {
	Width    int
	Height   int
	Settings []Setting
	CatIndex int
	SetIndex int

	state  *commit.State
	styles Styles
}

type Setting struct {
	Name     string
	Category string
}

const (
	defaultWidth  = 20
	defaultHeight = 20

	defaultCatSpacer = " "
	defaultCatPrompt = "❯"

	defaultSetSpacer = "  "
	defaultSetJoiner = "│ "
	defaultSetPrompt = "└▸"
)

func New(state *commit.State) Model {
	return Model{
		Width:  defaultWidth,
		Height: defaultHeight,
		styles: defaultStyles(state.Theme),
		state:  state,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg.(type) {
	case colour.Msg:
		m.styles = defaultStyles(m.state.Theme)
	}

	return m, nil
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		MaxWidth(m.Width).
		MaxHeight(m.Height).
		Render(m.renderSection())
}

func (m Model) SelectedCategory() string {
	if m.CatIndex >= len(Categories(m.Settings)) {
		return ""
	}

	return Categories(m.Settings)[m.CatIndex]
}

func (m Model) SelectedSetting() string {
	ss := Settings(m.SelectedCategory(), m.Settings)

	if len(ss) <= 1 {
		return ""
	}

	return ss[m.SetIndex]
}

func (m *Model) Reset() {
	m.CatIndex = 0
	m.SetIndex = 0
}

func (m *Model) Next() {
	// List of categories
	cat := Categories(m.Settings)
	// List of current category settings
	set := Settings(cat[m.CatIndex], m.Settings)

	lastCategory := len(cat) - 1
	lastSetting := len(set) - 1

	switch {
	// Index at end of setting but not on the last category
	case m.SetIndex >= lastSetting && !(m.CatIndex >= lastCategory):
		m.CatIndex++
		m.SetIndex = 0

	// Index at last category and last setting
	case m.SetIndex >= lastSetting && m.CatIndex >= lastCategory:
	default:
		m.SetIndex++
	}
}

func (m *Model) Previous() {
	// List of categories
	cat := Categories(m.Settings)

	switch {
	// Index at first category and first setting
	case m.CatIndex <= 0 && m.SetIndex <= 0:

	// Index at beginning of setting but not on the firat category
	case m.CatIndex != 0 && m.SetIndex <= 0:
		m.CatIndex--
		prevSet := Settings(cat[m.CatIndex], m.Settings)
		m.SetIndex = len(prevSet) - 1

	default:
		m.SetIndex--
	}
}

func (m Model) renderSection() string {
	var str []string

	for idx, c := range Categories(m.Settings) {
		selected := idx == m.CatIndex

		cp := m.styles.categorySpacer
		var cat string

		switch selected {
		case true:
			cp = m.styles.categoryPrompt
			cat = fmt.Sprintf("%v %v", cp, m.styles.categorySelected.Render(c))
		default:
			cat = fmt.Sprintf("%v %v", cp, m.styles.setting.Render(c))
		}

		str = append(str, cat)

		rc := m.renderCategory(c, selected)
		if len(rc) <= 1 {
			str = append(str, "")
			continue
		}

		str = append(str, rc)
	}

	return strings.Join(str, "\n")
}

func (m Model) renderCategory(cat string, selected bool) string {
	var str []string

	for idx, s := range Settings(cat, m.Settings) {
		var ss string

		sp := m.styles.settingSpacer

		switch {
		case idx < m.SetIndex && selected:
			sp = m.styles.settingJoiner
			ss = m.styles.setting.Render(s)
		case idx == m.SetIndex && selected:
			sp = m.styles.settingPrompt
			ss = m.styles.settingSelected.Render(s)
		default:
			ss = m.styles.setting.Render(s)
		}

		str = append(str, fmt.Sprintf("  %v%v", sp, ss))
	}

	return strings.Join(str, "\n") + "\n"
}

func Categories(sets []Setting) []string {
	var cat []string

	for _, v := range sets {
		if !slices.Contains(cat, v.Category) {
			cat = append(cat, v.Category)
		}
	}

	return cat
}

func Settings(cat string, sets []Setting) []string {
	var set []string

	for _, v := range sets {
		if cat != v.Category {
			continue
		}

		set = append(set, v.Name)
	}

	if len(set) <= 1 {
		return nil
	}

	return set
}

func ToModel(m tea.Model, c tea.Cmd) (Model, tea.Cmd) {
	return m.(Model), c
}
