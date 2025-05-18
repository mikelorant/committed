package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/terminal"
	"github.com/mikelorant/committed/internal/ui/body"
	"github.com/mikelorant/committed/internal/ui/colour"
	"github.com/mikelorant/committed/internal/ui/footer"
	"github.com/mikelorant/committed/internal/ui/header"
	"github.com/mikelorant/committed/internal/ui/help"
	"github.com/mikelorant/committed/internal/ui/info"
	"github.com/mikelorant/committed/internal/ui/message"
	"github.com/mikelorant/committed/internal/ui/option"
	"github.com/mikelorant/committed/internal/ui/status"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Request       *commit.Request
	Date          time.Time
	state         *commit.State
	focus         focus
	previousFocus focus
	models        Models
	quit          quit
	amend         bool
	file          bool
	signoff       bool
	err           error
	ready         bool
	writeConfig   bool
	currentSave   savedState
	previousSave  savedState
	emojiType     config.EmojiType
}

type Models struct {
	info    info.Model
	header  header.Model
	body    body.Model
	footer  footer.Model
	status  status.Model
	help    help.Model
	message message.Model
	option  option.Model
}

type savedState struct {
	amend   bool
	emoji   emoji.Emoji
	summary string
	body    string
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
	optionComponent
)

type quit int

const (
	unsetQuit quit = iota
	applyQuit
	cancelQuit
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

const (
	KeyAmend   = "å"
	KeyLoad    = "¬"
	KeySignoff = "ß"
	KeyTheme   = "†"
	KeyAuthor  = "¡"
	KeyEmoji   = "™"
	KeySummary = "£"
	KeyBody    = "¢"
	KeyHelp    = "˙"
	KeyOption  = "ø"
)

const dateTimeFormat = "Mon Jan 2 15:04:05 2006 -0700"

func New() Model {
	return Model{
		Date: time.Now(),
	}
}

func (m *Model) Configure(state *commit.State) {
	m.state = state
	m.defaults(state.Config)

	m.models = Models{
		info:   info.New(state),
		header: header.New(state),
		body:   body.New(state, bodyDefaultHeight),
		footer: footer.New(state),
		status: status.New(state),
		help:   help.New(state),
		option: option.New(state),
	}

	m.models.info.Date = m.Date.Format(dateTimeFormat)

	m.setSaves()
	m.restoreModel(m.currentSave)
	m.setCompatibility()
	m.configureOptions()

	if (m.state.Snapshot.Restore && m.setSave()) || m.file {
		m.resetCursor()
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

	if m.quit == applyQuit {
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

	if m.focus == optionComponent {
		return lipgloss.JoinVertical(lipgloss.Top,
			m.models.info.View(),
			m.models.option.View(),
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
	case "alt+1", KeyAuthor:
		if m.focus == authorComponent {
			return keyResponse{model: m, nilMsg: true}
		}
		m.focus = authorComponent
	case "alt+2", KeyEmoji:
		if m.focus == emojiComponent {
			return keyResponse{model: m, nilMsg: true}
		}
		m.focus = emojiComponent
	case "alt+3", KeySummary:
		if m.focus == summaryComponent {
			return keyResponse{model: m, nilMsg: true}
		}
		m.focus = summaryComponent
	case "alt+4", KeyBody:
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
	case "alt+enter", "alt+\\":
		if !m.validate() {
			break
		}

		m = m.commit(applyQuit)

		return keyResponse{model: m, cmd: tea.Quit, end: true}
	case "alt+a", KeyAmend:
		m.amend = !m.amend

		m.swapSave()

		m.models.header.CursorStartSummary()
		m.models.body.CursorStart()

		return keyResponse{model: m, end: false, nilMsg: true}
	case "alt+l", KeyLoad:
		if m.setSave() {
			m.models.header.CursorStartSummary()
			m.models.body.CursorStart()
		}

		return keyResponse{model: m, end: false, nilMsg: true}
	case "alt+s", KeySignoff:
		m.signoff = !m.signoff

		return keyResponse{model: m, end: false, nilMsg: true}
	case "alt+t", KeyTheme:
		m.state.Theme.Next()
		return keyResponse{model: m, cmd: colour.Update, end: true}
	case "ctrl+h", KeyHelp:
		if m.focus == helpComponent {
			m.focus = m.previousFocus
			break
		}
		m.previousFocus = m.focus
		m.focus = helpComponent
	case "ctrl+o", KeyOption:
		if m.focus == optionComponent {
			m.focus = m.previousFocus
			break
		}
		m.previousFocus = m.focus
		m.focus = optionComponent
	case "ctrl+w":
		m.state.Config = ToConfig(m.state.Config, m.models.option.GetPaneSets(), m.state.Theme)
		m.writeConfig = true
	case "esc":
		if m.focus == helpComponent || m.focus == optionComponent {
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
		m = m.commit(cancelQuit)

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
	m.models.footer.Signoff = m.signoff
	m.models.help.Blur()
	m.models.option.Blur()

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
	case optionComponent:
		m.models.status.Shortcuts = status.OptionShortcuts()
		m.models.option.Focus()
	}

	if m.signoff {
		m.models.body.Height -= footerSignoffHeight
	}

	return m
}

func (m Model) updateModels(msg tea.Msg) (Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 7)
	m.models.info, cmds[0] = info.ToModel(m.models.info.Update(msg))
	m.models.header, cmds[1] = header.ToModel(m.models.header.Update(msg))
	m.models.body, cmds[2] = body.ToModel(m.models.body.Update(msg))
	m.models.footer, cmds[3] = footer.ToModel(m.models.footer.Update(msg))
	m.models.status, cmds[4] = status.ToModel(m.models.status.Update(msg))
	m.models.help, cmds[5] = help.ToModel(m.models.help.Update(msg))

	if m.focus == optionComponent {
		m.models.option, cmds[6] = option.ToModel(m.models.option.Update(msg))
	}

	if !m.ready {
		m.ready = true
	}

	return m, tea.Batch(cmds...)
}

func (m Model) commit(q quit) Model {
	m.quit = q

	var emoji string

	switch m.emojiType {
	case config.EmojiTypeShortcode:
		emoji = m.models.header.Emoji.Shortcode
	default:
		emoji = m.models.header.Emoji.Character
	}

	if m.quit == applyQuit {
		m.models.message = message.New(message.State{
			Emoji:   emoji,
			Summary: m.models.header.Summary(),
			Body:    m.models.body.Value(),
			Footer:  m.models.footer.Value(),
			Theme:   m.state.Theme,
		})
	}

	m.Request = &commit.Request{
		Author:      m.models.info.Author,
		Emoji:       emoji,
		Summary:     m.models.header.Summary(),
		Body:        m.models.body.Value(),
		RawBody:     m.models.body.RawValue(),
		Footer:      m.models.footer.Value(),
		Amend:       m.amend,
		DryRun:      m.state.Options.DryRun,
		File:        m.file,
		MessageFile: m.state.Options.File.MessageFile,
	}

	if m.writeConfig {
		m.Request.Config = m.state.Config
		m.Request.Config.Update = true
	}

	if m.quit == applyQuit {
		m.Request.Apply = true

		return m
	}

	return m
}

func (m Model) validate() bool {
	staged := m.state.Repository.Worktree.IsStaged()
	summary := m.models.header.Summary()

	return (staged || m.amend) && (summary != "" || m.file)
}

func (m *Model) resetCursor() {
	m.models.header.CursorStartSummary()
	m.models.body.CursorStart()
}

func (m *Model) setCompatibility() {
	terminal.Set(m.state.Config.View.Compatibility)
}
