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
		output hook.Action
	}{
		{
			name:   "empty",
			output: hook.ActionUnset,
		},
		{
			name:   "install",
			input:  hook.Options{Install: true},
			output: hook.ActionInstall,
		},
		{
			name:   "uninstall",
			input:  hook.Options{Uninstall: true},
			output: hook.ActionUninstall,
		},
		{
			name:   "commit",
			input:  hook.Options{Commit: true},
			output: hook.ActionCommit,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := hook.New(tt.input)

			assert.Equal(t, tt.output, got.Action)
		})
	}
}

func TestDo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input hook.Action
		err   error
	}{
		{
			name:  "empty",
			input: hook.ActionUnset,
			err:   hook.ErrAction,
		},
		{
			name:  "install",
			input: hook.ActionInstall,
		},
		{
			name:  "uninstall",
			input: hook.ActionUninstall,
		},
		{
			name:  "commit",
			input: hook.ActionCommit,
			err:   hook.ErrAction,
		},
		{
			name: "nil",
			err:  hook.ErrAction,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hook := hook.Hook{
				Action: tt.input,
			}

			err := hook.Do()

			assert.ErrorIs(t, err, tt.err)
		})
	}
}
