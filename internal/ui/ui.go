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
	state         state
	previousState state
	models        Models
	request       *commit.Request
	quit          bool
	signoff       bool
	err           error
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

type state int

const (
	emptyComponent state = iota
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

func New(cfg commit.Config) (*commit.Request, error) {
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		fh, err := tea.LogToFile(logfilePath, "committed")
		if err != nil {
			return nil, fmt.Errorf("unable to log to file: %w", err)
		}
		defer fh.Close()
	}

	im := Model{
		state: emojiComponent,
		models: Models{
			info:   info.New(cfg),
			header: header.New(cfg),
			body:   body.New(cfg, bodyDefaultHeight),
			footer: footer.New(cfg),
			status: status.New(),
			help:   help.New(),
		},
	}

	p := tea.NewProgram(im)
	m, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("unable to run program: %w", err)
	}

	return m.(Model).request, nil
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
			if m.state == authorComponent {
				m.models.info, _ = info.ToModel(m.models.info.Update(msg))
				m.state = emojiComponent
				break
			}
			if m.state == emojiComponent {
				m.models.header, _ = header.ToModel(m.models.header.Update(msg))
				m.state = summaryComponent
				break
			}
			if m.state == summaryComponent {
				m.state = bodyComponent
			}
		case "alt+enter":
			m.request = &commit.Request{
				Author:  m.models.info.Author,
				Emoji:   m.models.header.Emoji.Shortcode,
				Summary: m.models.header.Summary(),
				Body:    m.models.body.Value(),
				Footer:  m.models.footer.Value(),
			}
			if m.validate() {
				m.quit = true
				m.message()
				return m, tea.Quit
			}
		case "alt+s":
			m.signoff = !m.signoff
		case "alt+t":
			cmd := theme.NextTint
			return m, cmd
		case "alt+/":
			if m.state == helpComponent {
				m.state = m.previousState
				break
			}
			m.previousState = m.state
			m.state = helpComponent
		case "esc":
			if m.state == helpComponent {
				m.state = m.previousState
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

	m.models.info.Blur()
	m.models.info.Expand = false
	m.models.header.Blur()
	m.models.header.Expand = false
	m.models.body.Blur()
	m.models.body.Height = bodyDefaultHeight
	m.models.footer.Author = m.models.info.Author
	m.models.help.Blur()

	switch m.state {
	case authorComponent:
		m.models.info.Focus()
		m.models.info.Expand = true
		m.models.body.Height = bodyAuthorHeight
		m.models.status.Previous = emptyName
		m.models.status.Next = emojiName
	case emojiComponent:
		m.models.header.Focus()
		m.models.header.SelectEmoji()
		m.models.header.Expand = true
		m.models.body.Height = bodyEmojiHeight
		m.models.status.Previous = authorName
		m.models.status.Next = summaryName
	case summaryComponent:
		m.models.header.Focus()
		m.models.header.SelectSummary()
		m.models.status.Previous = emojiName
		m.models.status.Next = bodyName
	case bodyComponent:
		m.models.body.Focus()
		m.models.status.Previous = summaryName
		m.models.status.Next = emptyName
	case helpComponent:
		m.models.status.Previous = emptyName
		m.models.status.Next = emptyName
		m.models.help.Focus()
	}

	if m.signoff {
		m.models.body.Height -= footerSignoffHeight
	}

	if m.signoff != m.models.footer.Signoff {
		m.models.footer.ToggleSignoff()
		m.models.body, _ = body.ToModel(m.models.body.Update(nil))
		m.models.footer, _ = footer.ToModel(m.models.footer.Update(nil))
		return m, nil
	}

	cmds := make([]tea.Cmd, 6)
	m.models.info, cmds[0] = info.ToModel(m.models.info.Update(msg))
	m.models.header, cmds[1] = header.ToModel(m.models.header.Update(msg))
	m.models.body, cmds[2] = body.ToModel(m.models.body.Update(msg))
	m.models.footer, cmds[3] = footer.ToModel(m.models.footer.Update(msg))
	m.models.status, cmds[4] = status.ToModel(m.models.status.Update(msg))
	m.models.help, cmds[5] = help.ToModel(m.models.help.Update(msg))

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("unable to render view: %s", m.err)
	}

	if m.quit {
		return lipgloss.JoinVertical(lipgloss.Top,
			m.models.info.View(),
			m.models.message.View(),
		)
	}

	if m.state == helpComponent {
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

func (m *Model) message() {
	mc := message.Config{
		Emoji:   m.models.header.Emoji,
		Summary: m.models.header.Summary(),
		Body:    m.models.body.Value(),
		Footer:  m.models.footer.Value(),
	}

	m.models.message = message.New(mc)
}

func (m Model) validate() bool {
	//nolint:gocritic
	switch {
	case m.request.Summary == "":
		return false
	}
	return true
}
