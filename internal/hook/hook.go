package hook

import (
	"errors"
	"io"

	"github.com/mikelorant/committed/internal/shell"
)

type Hook struct {
	Action Action
	Runner Runner
}

type Options struct {
	Install   bool
	Uninstall bool
	Commit    bool
}

type (
	Runner func(io.Writer, string, []string) error
)

type (
	Action int
)

var ErrAction = errors.New("invalid hook action")

const (
	ActionUnset Action = iota
	ActionInstall
	ActionUninstall
	ActionCommit
)

func New(opts Options) Hook {
	return Hook{
		Action: action(opts),
		Runner: shell.Run,
	}
}

func (h *Hook) Do() error {
	switch h.Action {
	case ActionInstall:
		return nil
	case ActionUninstall:
		return nil
	}

	return ErrAction
}

func action(opts Options) Action {
	switch {
	case opts.Install:
		return ActionInstall
	case opts.Uninstall:
		return ActionUninstall
	case opts.Commit:
		return ActionCommit
	default:
		return ActionUnset
	}
}
