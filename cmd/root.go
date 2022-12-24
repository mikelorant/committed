package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var opts commit.Options

	cmd := &cobra.Command{
		Use:   "committed",
		Short: "Committed is a WYSIWYG Git commit editor",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := commit.New(opts)
			if err != nil {
				log.Fatalf("unable to init commit: %v", err)
			}

			res, err := ui.New(c.Config)
			if err != nil {
				log.Fatalf("unable to init ui: %v", err)
			}

			if !res.Commit {
				return nil
			}

			if err := Commit(c, res); err != nil {
				log.Fatalf("unable to commit: %v", err)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Apply, "yes", "y", false, "Specify --yes to apply the commit")

	cc.Init(&cc.Config{
		RootCmd:         cmd,
		Headings:        cc.HiGreen + cc.Bold,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})

	return cmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func Commit(c *commit.Commit, res ui.Result) error {
	c.Author.Name = res.Author.Name
	c.Author.Email = res.Author.Email
	c.Emoji = res.Emoji
	c.Summary = res.Summary
	c.Body = strings.TrimSpace(res.Body)
	c.Footer = res.Footer

	if err := c.Create(); err != nil {
		return fmt.Errorf("unable to create commit: %w", err)
	}

	return nil
}
