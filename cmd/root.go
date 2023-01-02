package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/ui"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var hideApplyFlag bool

func NewRootCmd() *cobra.Command {
	var opts commit.Options
	var dryRun bool

	cmd := &cobra.Command{
		Use:         "committed",
		Short:       "Committed is a WYSIWYG Git commit editor",
		Version:     ReleaseVersion,
		Annotations: annotations(),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Apply = apply(opts.Apply, dryRun)

			c, err := commit.New(opts)
			switch {
			case err == nil:
			case errors.Is(err, repository.NotFoundError()):
				log.Fatalf("No git repository found.")
			default:
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

	cmd.AddCommand(NewVersionCmd())
	cmd.SetVersionTemplate(verTmpl)
	cmd.Flags().SortFlags = false
	cmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "Simulate applying a commit")
	cmd.Flags().BoolVarP(&opts.Amend, "amend", "a", false, "Replace the tip of the current branch by creating a new commit")

	if !hideApplyFlag {
		cmd.Flags().BoolVarP(&opts.Apply, "yes", "y", false, "Specify --yes to apply the commit")
	}

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

	if err := c.Apply(); err != nil {
		return fmt.Errorf("unable to create commit: %w", err)
	}

	return nil
}

func apply(a bool, dr bool) bool {
	if a {
		return true
	}

	if hideApplyFlag {
		a = true
	}

	if dr {
		a = false
	}

	return a
}
