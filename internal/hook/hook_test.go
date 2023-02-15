package hook_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/hook"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  hook.Options
		output hook.Hook
	}{
		{
			name:   "empty",
			output: hook.Hook{},
		},
		{
			name:   "install",
			input:  hook.Options{Install: true},
			output: hook.Hook{Action: hook.ActionInstall},
		},
		{
			name:   "uninstall",
			input:  hook.Options{Uninstall: true},
			output: hook.Hook{Action: hook.ActionUninstall},
		},
		{
			name:   "commit",
			input:  hook.Options{Commit: true},
			output: hook.Hook{Action: hook.ActionCommit},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := hook.New(tt.input)

			assert.Equal(t, tt.output, got)
		})
	}
}

func TestDo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input hook.Hook
		err   error
	}{
		{
			name:  "empty",
			input: hook.Hook{},
			err:   hook.ErrAction,
		},
		{
			name:  "install",
			input: hook.Hook{Action: hook.ActionInstall},
		},
		{
			name:  "uninstall",
			input: hook.Hook{Action: hook.ActionUninstall},
		},
		{
			name:  "commit",
			input: hook.Hook{Action: hook.ActionCommit},
			err:   hook.ErrAction,
		},
		{
			name:  "nil",
			input: hook.Hook{},
			err:   hook.ErrAction,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.input.Do()

			assert.ErrorIs(t, err, tt.err)
		})
	}
}
