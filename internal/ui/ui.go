package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/body"
)

type MainModel struct {
	state  component
	models Models
	config commit.Config
	result Result
	err    error
}

type Models struct {
	info   InfoModel
	header HeaderModel
	body   body.Model
	footer FooterModel
	status StatusModel
}

type State struct {
	component component
	display   display
}

type Result struct {
	Commit  bool
	Name    string
	Email   string
	Emoji   string
	Summary string
	Body    string
	Footer  string
}

type (
	component int
	display   int
)

const (
	emptyComponent component = iota
	authorComponent
	emojiComponent
	summaryComponent
	bodyComponent
)

const (
	defaultDisplay display = iota
	compactDisplay
	expandedDisplay
)

func New(cfg commit.Config) (Result, error) {
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		fh, err := tea.LogToFile(logfilePath, "committed")
		if err != nil {
			return Result{}, fmt.Errorf("unable to log to file: %w", err)
		}
		defer fh.Close()
	}

	im := MainModel{
		state: emojiComponent,
		models: Models{
			info:   NewInfo(cfg),
			header: NewHeader(cfg),
			body:   body.New(cfg),
			footer: NewFooter(cfg),
			status: NewStatus(cfg),
		},
		config: cfg,
	}

	p := tea.NewProgram(im)
	m, err := p.Run()
	if err != nil {
		return Result{}, fmt.Errorf("unable to run program: %w", err)
	}

	return m.(MainModel).result, nil
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(
		m.models.info.Init(),
		m.models.header.Init(),
		m.models.body.Init(),
		m.models.footer.Init(),
		m.models.status.Init(),
	)
}

//nolint:ireturn
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "alt+1":
			if m.state == authorComponent {
				return m, nil
			}
			m.state = authorComponent
		case "alt+2":
			if m.state == emojiComponent {
				return m, nil
			}
			m.state = emojiComponent
		case "alt+3":
			if m.state == summaryComponent {
				return m, nil
			}
			m.state = summaryComponent
		case "alt+4":
			if m.state == bodyComponent {
				return m, nil
			}
			m.state = bodyComponent
		case "enter":
			if m.state == emojiComponent {
				m.models.header, _ = m.models.header.Update(msg)
				m.state = summaryComponent
				break
			}
			if m.state == summaryComponent {
				m.state = bodyComponent
			}
		case "alt+enter":
			m.result = Result{
				Commit:  true,
				Name:    m.models.info.Name(),
				Email:   m.models.info.Email(),
				Emoji:   m.models.header.EmojiShortCode(),
				Summary: m.models.header.Summary(),
				Body:    m.models.body.Value(),
			}
			if m.validate() {
				return m, tea.Quit
			}
		case "tab":
			switch m.state {
			case authorComponent:
				m.state = emojiComponent
			case emojiComponent:
				m.state = summaryComponent
			case summaryComponent:
				m.state = bodyComponent
			}
		case "shift+tab":
			switch m.state {
			case emojiComponent:
				m.state = authorComponent
			case summaryComponent:
				m.state = emojiComponent
			case bodyComponent:
				m.state = summaryComponent
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.models.info.state = State{}
	m.models.header.state = State{}
	m.models.body.Blur()
	m.models.body.Compact = false
	m.models.footer.state = State{}
	m.models.status.state = State{}

	switch m.state {
	case authorComponent:
		m.models.header.state.component = authorComponent
	case emojiComponent:
		m.models.header.state.component = emojiComponent
		m.models.header.state.display = expandedDisplay
		m.models.body.Compact = true
	case summaryComponent:
		m.models.header.state.component = summaryComponent
	case bodyComponent:
		m.models.body.Focus()
	}

	cmds := make([]tea.Cmd, 5)
	m.models.info, cmds[0] = m.models.info.Update(msg)
	m.models.header, cmds[1] = m.models.header.Update(msg)
	m.models.body, cmds[2] = m.models.body.Update(msg)
	m.models.footer, cmds[3] = m.models.footer.Update(msg)
	m.models.status, cmds[4] = m.models.status.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("unable to render view: %s", m.err)
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		m.models.info.View(),
		m.models.header.View(),
		m.models.body.View(),
		m.models.footer.View(),
		m.models.status.View(),
	)
}

func (m MainModel) validate() bool {
	//nolint:gocritic
	switch {
	case m.result.Summary == "":
		return false
	}
	return true
}
