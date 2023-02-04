package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/emoji"
	"github.com/mikelorant/committed/internal/theme"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func NewListCmd(w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List settings with profiles or IDs",
	}

	cmd.AddCommand(NewListThemesCmd(w))
	cmd.AddCommand(NewListEmojiProfilesCmd(w))

	return cmd
}

func NewListThemesCmd(w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "themes",
		Short: "List theme IDs",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			listThemes(w)
		},
	}

	return cmd
}

func NewListEmojiProfilesCmd(w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "emojis",
		Short: "List emoji profiles",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			listEmojiProfiles(w)
		},
	}

	return cmd
}

func listThemes(w io.Writer) {
	th := theme.New(config.ColourAdaptive)

	tbl := table.New("Name", "ID", "Palette")
	tbl.WithHeaderFormatter(header(th.Registry))
	tbl.WithWidthFunc(lipgloss.Width)
	tbl.WithWriter(w)

	for _, v := range th.List() {
		th.Set(v.ID())
		tbl.AddRow(v.DisplayName(), v.ID(), palette(th.Registry))
	}

	tbl.Print()
}

func listEmojiProfiles(w io.Writer) {
	em := emoji.New()
	th := theme.New(config.ColourAdaptive)

	tbl := table.New("Profile", "URL")
	tbl.WithHeaderFormatter(header(th.Registry))
	tbl.WithWidthFunc(lipgloss.Width)
	tbl.WithWriter(w)

	for _, v := range em.ListProfiles()[1:] {
		tbl.AddRow(emoji.ToString(v), emoji.ToURL(v))
	}

	tbl.Print()
}

func header(reg *tint.Registry) func(format string, vals ...interface{}) string {
	fg := reg.BrightBlack()
	if lipgloss.HasDarkBackground() {
		fg = reg.BrightWhite()
	}

	return func(format string, vals ...interface{}) string {
		var ss []any

		for _, v := range vals {
			style := lipgloss.NewStyle().Foreground(fg)
			ss = append(ss, style.Render(v.(string)))
		}

		return fmt.Sprintf(format, ss...)
	}
}

func palette(reg *tint.Registry) string {
	var colours []string

	pal := []lipgloss.TerminalColor{
		reg.Black(), reg.Red(), reg.Green(), reg.Yellow(),
		reg.Blue(), reg.Purple(), reg.Cyan(), reg.White(),
		reg.BrightBlack(), reg.BrightRed(), reg.BrightGreen(), reg.BrightYellow(),
		reg.BrightBlue(), reg.BrightPurple(), reg.BrightCyan(), reg.BrightWhite(),
	}

	for _, v := range pal {
		colours = append(colours, lipgloss.NewStyle().Background(v).Render(" "))
	}

	return strings.Join(colours, "")
}
