package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/repository"
	"github.com/mikelorant/committed/internal/ui"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var defaultDryRun = true

func NewRootCmd() *cobra.Command {
	var opts commit.Options

	cmd := &cobra.Command{
		Use:         "committed",
		Short:       "Committed is a WYSIWYG Git commit editor",
		Version:     ReleaseVersion,
		Annotations: annotations(),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := commit.New(opts)
			switch {
			case err == nil:
			case errors.Is(err, repository.NotFoundError()):
				log.Fatalf("No git repository found.")
			default:
				log.Fatalf("unable to init commit: %v", err)
			}

			req, err := ui.New(c.Config).Start()
			if err != nil {
				log.Fatalf("unable to init ui: %v", err)
			}

			if err := c.Apply(req); err != nil {
				log.Fatalf("unable to commit: %v", err)
			}

			return nil
		},
	}

	cmd.AddCommand(NewVersionCmd())
	cmd.SetVersionTemplate(verTmpl)
	cmd.Flags().SortFlags = false
	cmd.Flags().BoolVarP(&opts.DryRun, "dry-run", "", defaultDryRun, "Simulate applying a commit")
	cmd.Flags().BoolVarP(&opts.Amend, "amend", "a", false, "Replace the tip of the current branch by creating a new commit")

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
