package terminal

import (
	"github.com/mikelorant/committed/internal/config"

	"github.com/rivo/uniseg"
)

type graphemes struct {
	codepoints []rune
	width      int
}

func Set(c config.Compatibility) {
	uniseg.GraphemeClusterWidthOverrides = overrideGraphemeClusterWidth(c)
}

func Clear() {
	uniseg.GraphemeClusterWidthOverrides = nil
}

func overrideGraphemeClusterWidth(c config.Compatibility) map[string]int {
	gs := make([]graphemes, 0)

	switch c {
	case config.CompatibilityUnicode9:
		gs = append(gs, overrideVS16()...)
	default:
	}

	return overrides(gs)
}

// Grapheme clusters using variant selector 16
// had their widths changed as part of Unicode 14.
// Unicode < 14 = 1.
// Unicode >= 14 = 2.
// Required for:
// - macOS Terminal (2.12.7)
// - iTerm2 (3.4.23)
// - VSCode (1.87.0)
// - Alacritty (0.13.1)
// - WezTerm (20240203)
func overrideVS16() []graphemes {
	return []graphemes{
		{codepoints: []rune{0x203c, 0xfe0f}, width: 1}, // ‼️
		{codepoints: []rune{0x21a9, 0xfe0f}, width: 1}, // ↩️
		{codepoints: []rune{0x2601, 0xfe0f}, width: 1}, // ☁️
		{codepoints: []rune{0x267b, 0xfe0f}, width: 1}, // ♻️
		{codepoints: []rune{0x2697, 0xfe0f}, width: 1}, // ⚗️
		{codepoints: []rune{0x2699, 0xfe0f}, width: 1}, // ⚙️
		{codepoints: []rune{0x26b0, 0xfe0f}, width: 1}, // ⚰️
		{codepoints: []rune{0x270f, 0xfe0f}, width: 1}, // ✏️
		{codepoints: []rune{0x2b06, 0xfe0f}, width: 1}, // ⬆️
		{codepoints: []rune{0x2b07, 0xfe0f}, width: 1}, // ⬇️
	}
}

func overrides(gs []graphemes) map[string]int {
	overrides := make(map[string]int, len(gs))
	for _, g := range gs {
		key := string(g.codepoints)
		overrides[key] = g.width
	}

	return overrides
}
