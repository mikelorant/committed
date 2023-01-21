package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui"

	"github.com/go-git/go-git/v5"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

type Commiter interface {
	Configure(opts commit.Options) (*commit.Config, error)
	Apply(req *commit.Request) error
}

type UIer interface {
	Configure(cfg *commit.Config)
	Start() (*commit.Request, error)
}

type Logger interface {
	Fatalf(format string, v ...any)
}

type App struct {
	Commiter Commiter
	UIer     UIer
	Logger   Logger

	req  *commit.Request
	opts commit.Options
}

var defaultDryRun = true

func NewRootCmd(a App) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "committed",
		Short:       "Committed is a WYSIWYG Git commit editor",
		Version:     ReleaseVersion,
		Annotations: annotations(),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return a.configure(a.opts)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.start()
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return a.apply()
		},
	}

	cmd.AddCommand(NewVersionCmd())
	cmd.SetVersionTemplate(verTmpl)
	cmd.Flags().SortFlags = false
	cmd.Flags().BoolVarP(&a.opts.DryRun, "dry-run", "", defaultDryRun, "Simulate applying a commit")
	cmd.Flags().BoolVarP(&a.opts.Amend, "amend", "a", false, "Replace the tip of the current branch by creating a new commit")

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
	if err := NewRootCmd(NewApp()).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func NewApp() App {
	c := commit.New()
	u := ui.New()
	l := log.Default()

	return App{
		Commiter: &c,
		UIer:     &u,
		Logger:   l,
	}
}

func (a *App) configure(opts commit.Options) error {
	cfg, err := a.Commiter.Configure(opts)
	switch {
	case err == nil:
	case errors.Is(err, git.ErrRepositoryNotExists):
		a.Logger.Fatalf("No git repository found.")
		return err
	default:
		a.Logger.Fatalf("unable to init commit: %v", err)
		return err
	}

	a.UIer.Configure(cfg)

	return nil
}

func (a *App) start() error {
	r, err := a.UIer.Start()
	if err != nil {
		a.Logger.Fatalf("unable to start ui: %v", err)
		return err
	}
	a.req = r

	return nil
}

func (a *App) apply() error {
	if err := a.Commiter.Apply(a.req); err != nil {
		a.Logger.Fatalf("unable to apply commit: %v", err)
		return err
	}

	return nil
}
