package ui

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/ui/body"
	"github.com/mikelorant/committed/internal/ui/colour"
	"github.com/mikelorant/committed/internal/ui/footer"
	"github.com/mikelorant/committed/internal/ui/header"
	"github.com/mikelorant/committed/internal/ui/help"
	"github.com/mikelorant/committed/internal/ui/info"
	"github.com/mikelorant/committed/internal/ui/message"
	"github.com/mikelorant/committed/internal/ui/status"
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
	signoff       bool
	err           error
	ready         bool
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
)

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
	}

	dateTimeFormat := "Mon Jan 2 15:04:05 2006 -0700"
	m.models.info.Date = m.Date.Format(dateTimeFormat)

	m.amend = state.Options.Amend

	switch m.amend {
	case true:
		m.currentSave = defaultAmendSave(state)
		m.previousSave = savedState{}
	case false:
		m.currentSave = savedState{}
		m.previousSave = defaultAmendSave(state)
	}

	m.restoreModel(m.currentSave)

	if m.state.Snapshot.Restore {
		if m.setSave() {
			m.models.header.CursorStartSummary()
			m.models.body.CursorStart()
		}
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

func (m Model) commit(q quit) Model {
	m.quit = q

	var emoji string

	switch m.emojiType {
	case config.EmojiTypeCharacter:
		emoji = m.models.header.Emoji.Character
	default:
		emoji = m.models.header.Emoji.Shortcode
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
		Author:  m.models.info.Author,
		Emoji:   emoji,
		Summary: m.models.header.Summary(),
		Body:    m.models.body.Value(),
		RawBody: m.models.body.RawValue(),
		Footer:  m.models.footer.Value(),
		Amend:   m.amend,
	}

	if m.quit == applyQuit {
		m.Request.Apply = true

		return m
	}

	return m
}

func (m Model) validate() bool {
	return m.state.Repository.Worktree.IsStaged() && m.models.header.Summary() != ""
}

func (m *Model) restoreModel(save savedState) {
	m.models.header.Amend = save.amend
	m.models.header.Emoji = save.emoji
	m.models.header.SetSummary(save.summary)
	m.models.body.SetValue(save.body)
}

func (m *Model) backupModel() savedState {
	var save savedState

	save.amend = m.models.header.Amend
	save.emoji = m.models.header.Emoji
	save.summary = m.models.header.Summary()
	save.body = m.models.body.RawValue()

	return save
}

func (m *Model) setSave() bool {
	save := m.snapshotToSave()

	switch {
	case m.currentSave.amend && save.amend:
		m.loadSave(save)
		return true
	case m.previousSave.amend && save.amend:
		m.swapSave()
		m.loadSave(save)
		return true
	case (save.body != "" || save.emoji.Name != "" || save.summary != "") && !m.currentSave.amend:
		m.loadSave(save)
		return true
	case (save.body != "" || save.emoji.Name != "" || save.summary != "") && !m.previousSave.amend:
		m.swapSave()
		m.loadSave(save)
		return true
	}

	return false
}

func (m *Model) swapSave() {
	m.currentSave = m.backupModel()

	m.models.header.ResetSummary()
	m.models.body.Reset()

	m.currentSave, m.previousSave = m.previousSave, m.currentSave

	m.restoreModel(m.currentSave)
}

func (m *Model) loadSave(st savedState) {
	m.models.header.ResetSummary()
	m.models.body.Reset()

	m.restoreModel(st)
}
