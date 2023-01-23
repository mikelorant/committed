package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/body"
	"github.com/mikelorant/committed/internal/ui/footer"
	"github.com/mikelorant/committed/internal/ui/header"
	"github.com/mikelorant/committed/internal/ui/help"
	"github.com/mikelorant/committed/internal/ui/info"
	"github.com/mikelorant/committed/internal/ui/message"
	"github.com/mikelorant/committed/internal/ui/status"
	"github.com/mikelorant/committed/internal/ui/theme"
)

type Model struct {
	Request       *commit.Request
	focus         focus
	previousFocus focus
	models        Models
	quit          bool
	signoff       bool
	err           error
	ready         bool
}

type Models struct {
	info    info.Model
	header  header.Model
	body    body.Model
	footer  footer.Model
	status  status.Model
	help    help.Model
	message message.Model
}

type keyResponse struct {
	model  Model
	cmd    tea.Cmd
	end    bool
	nilMsg bool
}

type focus int

const (
	emptyComponent focus = iota
	authorComponent
	emojiComponent
	summaryComponent
	bodyComponent
	helpComponent
)

const (
	bodyDefaultHeight = 19
	bodyAuthorHeight  = 12
	bodyEmojiHeight   = 6

	footerSignoffHeight = 2
)

const (
	emptyName   = ""
	authorName  = "Author"
	emojiName   = "Emoji"
	summaryName = "Summary"
	bodyName    = "Body"
)

func New() Model {
	return Model{
		focus: emojiComponent,
	}
}

func (m *Model) Configure(state *commit.State) {
	m.models = Models{
		info:   info.New(state),
		header: header.New(state),
		body:   body.New(state, bodyDefaultHeight),
		footer: footer.New(state),
		status: status.New(),
		help:   help.New(),
	}
}

func (m Model) Start() (*commit.Request, error) {
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		fh, err := tea.LogToFile(logfilePath, "committed")
		if err != nil {
			return nil, fmt.Errorf("unable to log to file: %w", err)
		}
		defer fh.Close()
	}

	p := tea.NewProgram(m)
	r, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("unable to run program: %w", err)
	}

	return r.(Model).Request, nil
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.models.info.Init(),
		m.models.header.Init(),
		m.models.body.Init(),
		m.models.footer.Init(),
		m.models.status.Init(),
		m.models.help.Init(),
	)
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msgType := msg.(type) {
	case tea.KeyMsg:
		resp := m.onKeyPress(msgType)
		switch {
		case resp.end:
			return resp.model, resp.cmd
		case resp.nilMsg:
			msg = nil
		}

		m = resp.model
	}

	m = m.resetModels()
	m = m.setModels()

	return m.updateModels(msg)
}

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("unable to render view: %s", m.err)
	}

	if !m.ready {
		return ""
	}

	if m.quit {
		return lipgloss.JoinVertical(lipgloss.Top,
			m.models.info.View(),
			m.models.message.View(),
		)
	}

	if m.focus == helpComponent {
		return lipgloss.JoinVertical(lipgloss.Top,
			m.models.info.View(),
			m.models.help.View(),
			m.models.status.View(),
		)
	}

	if !m.models.footer.Signoff {
		return lipgloss.JoinVertical(lipgloss.Top,
			m.models.info.View(),
			m.models.header.View(),
			m.models.body.View(),
			m.models.status.View(),
		)
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		m.models.info.View(),
		m.models.header.View(),
		m.models.body.View(),
		m.models.footer.View(),
		m.models.status.View(),
	)
}

func (m Model) onKeyPress(msg tea.KeyMsg) keyResponse {
	switch msg.String() {
	case "alt+1":
		if m.focus == authorComponent {
			return keyResponse{model: m, nilMsg: true}
		}
		m.focus = authorComponent
	case "alt+2":
		if m.focus == emojiComponent {
			return keyResponse{model: m, nilMsg: true}
		}
		m.focus = emojiComponent
	case "alt+3":
		if m.focus == summaryComponent {
			return keyResponse{model: m, nilMsg: true}
		}
		m.focus = summaryComponent
	case "alt+4":
		if m.focus == bodyComponent {
			return keyResponse{model: m, nilMsg: true}
		}
		m.focus = bodyComponent
	case "enter":
		switch m.focus {
		case authorComponent:
			m.models.info, _ = info.ToModel(m.models.info.Update(msg))
			m.focus = emojiComponent
		case emojiComponent:
			m.models.header, _ = header.ToModel(m.models.header.Update(msg))
			m.focus = summaryComponent
		case summaryComponent:
			m.focus = bodyComponent
		}
	case "alt+enter":
		if !m.validate() {
			break
		}

		m = m.commit()

		return keyResponse{model: m, cmd: tea.Quit, end: true}
	case "alt+s":
		m.signoff = !m.signoff
		m.models.footer.ToggleSignoff()

		return keyResponse{model: m, end: false, nilMsg: true}
	case "alt+t":
		return keyResponse{model: m, cmd: theme.NextTint, end: true}
	case "alt+/":
		if m.focus == helpComponent {
			m.focus = m.previousFocus
			break
		}
		m.previousFocus = m.focus
		m.focus = helpComponent
	case "esc":
		if m.focus == helpComponent {
			m.focus = m.previousFocus
		}
	case "tab":
		switch m.focus {
		case authorComponent:
			m.focus = emojiComponent
		case emojiComponent:
			m.focus = summaryComponent
		case summaryComponent:
			m.focus = bodyComponent
		}
	case "shift+tab":
		switch m.focus {
		case emojiComponent:
			m.focus = authorComponent
		case summaryComponent:
			m.focus = emojiComponent
		case bodyComponent:
			m.focus = summaryComponent
		}
	case "ctrl+c":
		return keyResponse{model: m, cmd: tea.Quit, end: true}
	}

	return keyResponse{model: m}
}

func (m Model) resetModels() Model {
	m.models.info.Blur()
	m.models.info.Expand = false
	m.models.header.Blur()
	m.models.header.Expand = false
	m.models.body.Blur()
	m.models.body.Height = bodyDefaultHeight
	m.models.footer.Author = m.models.info.Author
	m.models.help.Blur()

	return m
}

func (m Model) setModels() Model {
	switch m.focus {
	case authorComponent:
		m.models.info.Focus()
		m.models.info.Expand = true
		m.models.body.Height = bodyAuthorHeight
		m.models.status.Shortcuts = status.GlobalShortcuts(emojiName, emptyName)
	case emojiComponent:
		m.models.header.Focus()
		m.models.header.SelectEmoji()
		m.models.header.Expand = true
		m.models.body.Height = bodyEmojiHeight
		m.models.status.Shortcuts = status.GlobalShortcuts(summaryName, authorName)
	case summaryComponent:
		m.models.header.Focus()
		m.models.header.SelectSummary()
		m.models.status.Shortcuts = status.GlobalShortcuts(bodyName, emojiName)
	case bodyComponent:
		m.models.body.Focus()
		m.models.status.Shortcuts = status.GlobalShortcuts(emptyName, summaryName)
	case helpComponent:
		m.models.status.Shortcuts = status.HelpShortcuts()
		m.models.help.Focus()
	}

	if m.signoff {
		m.models.body.Height -= footerSignoffHeight
	}

	return m
}

func (m Model) updateModels(msg tea.Msg) (Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 6)
	m.models.info, cmds[0] = info.ToModel(m.models.info.Update(msg))
	m.models.header, cmds[1] = header.ToModel(m.models.header.Update(msg))
	m.models.body, cmds[2] = body.ToModel(m.models.body.Update(msg))
	m.models.footer, cmds[3] = footer.ToModel(m.models.footer.Update(msg))
	m.models.status, cmds[4] = status.ToModel(m.models.status.Update(msg))
	m.models.help, cmds[5] = help.ToModel(m.models.help.Update(msg))

	if !m.ready {
		m.ready = true
	}

	return m, tea.Batch(cmds...)
}

func (m Model) commit() Model {
	m.quit = true

	m.models.message = message.New(message.State{
		Emoji:   m.models.header.Emoji,
		Summary: m.models.header.Summary(),
		Body:    m.models.body.Value(),
		Footer:  m.models.footer.Value(),
	})

	m.Request = &commit.Request{
		Author:  m.models.info.Author,
		Emoji:   m.models.header.Emoji.Shortcode,
		Summary: m.models.header.Summary(),
		Body:    m.models.body.Value(),
		Footer:  m.models.footer.Value(),
	}

	return m
}

func (m Model) validate() bool {
	return m.models.header.Summary() != ""
}
