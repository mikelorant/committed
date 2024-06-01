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
			name:  "unicode14 empty",
			width: 0,
		},
		{
			name:  "unicode14 simple",
			value: "test",
			width: 4,
		},
		{
			name:  "unicode14 single emoji",
			value: "❤️",
			width: 2,
		},
		{
			name:  "unicode14 single vs16 emoji",
			value: "⬆️",
			width: 2,
		},
		{
			name:  "unicode14 multiple emoji",
			value: "❤️❤️",
			width: 4,
		},
		{
			name:  "unicode14 multiple vs16 emoji",
			value: "⬆️⬆️",
			width: 4,
		},
		{
			name:  "unicode14 multiple mixed emoji",
			value: "❤️⬆️",
			width: 4,
		},
		{
			name:   "unicode9 empty",
			compat: config.CompatibilityUnicode9,
			width:  0,
		},
		{
			name:   "unicode9 simple",
			compat: config.CompatibilityUnicode9,
			value:  "test",
			width:  4,
		},
		{
			name:   "unicode9 single emoji",
			compat: config.CompatibilityUnicode9,
			value:  "❤️",
			width:  2,
		},
		{
			name:   "unicode9 single vs16 emoji",
			compat: config.CompatibilityUnicode9,
			value:  "⬆️",
			width:  1,
		},
		{
			name:   "unicode9 multiple emoji",
			compat: config.CompatibilityUnicode9,
			value:  "❤️❤️",
			width:  4,
		},
		{
			name:   "unicode9 multiple vs16 emoji",
			compat: config.CompatibilityUnicode9,
			value:  "⬆️⬆️",
			width:  2,
		},
		{
			name:   "unicode9 multiple mixed emoji",
			compat: config.CompatibilityUnicode9,
			value:  "❤️⬆️",
			width:  3,
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
			name:    "unicode9 without override",
			value:   "❤️",
			width:   2,
			orWidth: 2,
		},
		{
			name:    "unicode9 with override",
			value:   "⬆️",
			width:   2,
			orWidth: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			width := uniseg.StringWidth(tt.value)
			assert.Equal(t, tt.width, width)

			terminal.Set(config.CompatibilityUnicode9)

			width = uniseg.StringWidth(tt.value)
			assert.Equal(t, tt.orWidth, width)

			terminal.Clear()

			width = uniseg.StringWidth(tt.value)
			assert.Equal(t, tt.width, width)
		})
	}
}
