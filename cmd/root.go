package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/hook"
	"github.com/mikelorant/committed/internal/ui"

	"github.com/go-git/go-git/v5"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

type Commiter interface {
	Configure(opts commit.Options) (*commit.State, error)
	Apply(req *commit.Request) error
}

type UIer interface {
	Configure(cfg *commit.State)
	Start() (*commit.Request, error)
}

type Logger interface {
	Fatalf(format string, v ...any)
}

type Hooker interface {
	Do(opts hook.Options) error
}

type App struct {
	Commiter Commiter
	UIer     UIer
	Logger   Logger
	Writer   io.Writer
	Hooker   Hooker

	req  *commit.Request
	opts commit.Options
	hook bool
}

type Options struct {
	Hook bool
}

func NewRootCmd(a App) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "committed",
		Short:       "Committed is a WYSIWYG Git commit editor",
		Version:     version,
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

	var (
		defaultDryRun       = isDryRun()
		defaultConfigFile   = "$HOME/.config/committed/config.yaml"
		defaultSnapshotFile = "$HOME/.local/state/committed/snapshot.yaml"
	)

	cmd.AddCommand(NewVersionCmd())
	cmd.AddCommand(NewListCmd(a.Writer))
	cmd.AddCommand(NewHookCmd(a))
	cmd.SetVersionTemplate(verTmpl)
	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&a.opts.ConfigFile, "config", "", defaultConfigFile, "Config file location")
	cmd.Flags().StringVarP(&a.opts.SnapshotFile, "snapshot", "", defaultSnapshotFile, "Snapshot file location")
	cmd.Flags().BoolVarP(&a.opts.DryRun, "dry-run", "", defaultDryRun, "Simulate applying a commit")
	cmd.Flags().BoolVarP(&a.opts.Amend, "amend", "a", false, "Replace the tip of the current branch by creating a new commit")
	cmd.Flags().StringVarP(&a.opts.File.MessageFile, "editor", "", "", "")
	cmd.Flags().BoolVarP(&a.hook, "hook", "", false, "")
	cmd.Flags().StringVarP(&a.opts.File.MessageFile, "message-file", "", "", "")
	cmd.Flags().StringVarP(&a.opts.File.Source, "source", "", "", "")
	cmd.Flags().StringVarP(&a.opts.File.SHA, "sha", "", "", "")
	cmd.Flags().MarkHidden("editor")
	cmd.Flags().MarkHidden("hook")
	cmd.Flags().MarkHidden("message-file")
	cmd.Flags().MarkHidden("source")
	cmd.Flags().MarkHidden("sha")

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
	h := hook.New()
	l := log.Default()
	u := ui.New()
	w := os.Stdout

	return App{
		Commiter: &c,
		Hooker:   &h,
		Logger:   l,
		UIer:     &u,
		Writer:   w,
	}
}

func (a *App) configure(opts commit.Options) error {
	opts.Mode = a.mode()

	state, err := a.Commiter.Configure(opts)
	switch {
	case err == nil:
	case errors.Is(err, git.ErrRepositoryNotExists):
		a.Logger.Fatalf("No git repository found.")
		return err
	default:
		a.Logger.Fatalf("unable to init commit: %v", err)
		return err
	}

	a.UIer.Configure(state)

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

func (a *App) mode() commit.Mode {
	switch {
	case !a.hook && a.opts.File.MessageFile != "":
		return commit.ModeEditor
	case a.hook:
		return commit.ModeHook
	default:
		return commit.ModeCommit
	}
}
