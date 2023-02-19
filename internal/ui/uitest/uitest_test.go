package uitest_test

import (
	"fmt"
	"testing"

	"github.com/mikelorant/committed/internal/ui/uitest"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

type MockModel struct {
	keys string
}

func (m MockModel) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m MockModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.keys += msg.(tea.KeyMsg).String()

	return m, nil
}

func (m MockModel) View() string {
	return m.keys
}

func TestStripString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "text",
			input:  "test",
			output: "test",
		},
		{
			name:   "text_multiline",
			input:  "test",
			output: "test",
		},
		{
			name: "ansi",
			input: lipgloss.NewStyle().
				Foreground(lipgloss.Color("9")).
				Render("test"),
			output: "test",
		},
		{
			name: "ansi_mixed",
			input: fmt.Sprintf("before %v after", lipgloss.NewStyle().
				Foreground(lipgloss.Color("9")).
				Render("test")),
			output: "before test after",
		},
		{
			name: "ansi_mixed_multiline",
			input: fmt.Sprintf("before\n%v\nafter", lipgloss.NewStyle().
				Foreground(lipgloss.Color("9")).
				Render("test")),
			output: "before\ntest\nafter",
		},
		{
			name:   "trailing_whitespace",
			input:  "test    ",
			output: "test",
		},
		{
			name: "trailing_ansi_whitespace",
			input: fmt.Sprintf("before %v    ", lipgloss.NewStyle().
				Foreground(lipgloss.Color("9")).
				Render("test")),
			output: "before test",
		},
		{
			name:   "trailing_multiline_whitespace",
			input:  "test    \ntest",
			output: "test\ntest",
		},
		{
			name:   "trailing_multiline_whitespace_multiple",
			input:  "test    \ntest    ",
			output: "test\ntest",
		},
		{
			name:   "empty",
			input:  "",
			output: "",
		},
		{
			name:   "empty_multiline",
			input:  "test\n\n\ntest",
			output: "test\n\n\ntest",
		},
		{
			name:   "tab",
			input:  "test\t",
			output: "test",
		},
		{
			name:   "emoji",
			input:  "test🎨",
			output: "test🎨",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.output, uitest.StripString(tt.input))
		})
	}
}

func TestKeyPress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		key   []rune
		value string
	}{
		{
			name:  "text",
			key:   []rune{'t', 'e', 's', 't'},
			value: "test",
		},
		{
			name:  "empty",
			key:   []rune{},
			value: "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var m tea.Model
			m = MockModel{}

			for _, k := range tt.key {
				m, _ = m.Update(uitest.KeyPress(k))
			}

			assert.Equal(t, tt.value, m.View())
		})
	}
}

func TestSendString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		res  string
	}{
		{
			name: "text",
			str:  "test",
			res:  "test",
		},
		{
			name: "empty",
			str:  "",
			res:  "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var m tea.Model
			m = MockModel{}

			m = uitest.SendString(m, tt.str)

			assert.Equal(t, tt.res, m.View())
		})
	}
}
