package terminal_test

import (
	"testing"

	"github.com/mikelorant/committed/internal/config"
	"github.com/mikelorant/committed/internal/terminal"

	"github.com/rivo/uniseg"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	tests := []struct {
		name   string
		compat config.Compatibility
		value  string
		width  int
	}{
		{
			name:  "default empty",
			width: 0,
		},
		{
			name:  "default simple",
			value: "test",
			width: 4,
		},
		{
			name:  "default single emoji without override",
			value: "❤️",
			width: 2,
		},
		{
			name:  "default single emoji with override",
			value: "⬆️",
			width: 1,
		},
		{
			name:  "default multiple emojis without override",
			value: "❤️❤️",
			width: 4,
		},
		{
			name:  "default multiple emojis with override",
			value: "⬆️⬆️",
			width: 2,
		},
		{
			name:  "default mixed emojis",
			value: "⬆️❤️",
			width: 3,
		},
		{
			name:  "default multiple mixed emojis",
			value: "⬆️❤️❤️⬆️",
			width: 6,
		},
		{
			name:   "ttyd empty",
			compat: config.CompatibilityTtyd,
			value:  "",
			width:  0,
		},
		{
			name:   "ttyd simple",
			compat: config.CompatibilityTtyd,
			value:  "test",
			width:  4,
		},
		{
			name:   "ttyd multiple emojis",
			compat: config.CompatibilityTtyd,
			value:  "⬆️❤️",
			width:  4,
		},
		{
			name:   "kitty empty",
			compat: config.CompatibilityKitty,
			value:  "",
			width:  0,
		},
		{
			name:   "ttyd simple",
			compat: config.CompatibilityKitty,
			value:  "test",
			width:  4,
		},
		{
			name:   "ttyd multiple emojis",
			compat: config.CompatibilityKitty,
			value:  "⬆️❤️",
			width:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminal.Set(tt.compat)

			width := uniseg.StringWidth(tt.value)
			assert.Equal(t, tt.width, width)

			terminal.Clear()
		})
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		width   int
		orWidth int
	}{
		{
			name:    "default without override",
			value:   "❤️",
			width:   2,
			orWidth: 2,
		},
		{
			name:    "default with override",
			value:   "⬆️",
			width:   2,
			orWidth: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			width := uniseg.StringWidth(tt.value)
			assert.Equal(t, tt.width, width)

			terminal.Set(config.CompatibilityDefault)

			width = uniseg.StringWidth(tt.value)
			assert.Equal(t, tt.orWidth, width)

			terminal.Clear()

			width = uniseg.StringWidth(tt.value)
			assert.Equal(t, tt.width, width)
		})
	}
}
