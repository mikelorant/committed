package shortcut

import (
	"fmt"
	"strings"

	"github.com/mikelorant/committed/internal/commit"

	"github.com/charmbracelet/lipgloss"
)

type KeyBinding struct {
	Modifier int
	Key      string
	Label    string
}

type Modifier struct {
	Modifier int
	Label    string
	Align    int
}

type keySet struct {
	align       int
	keybindings []KeyBinding
	modifiers   []Modifier
	target      [][]string
	decorate    bool
	height      int
	width       int
	state       *commit.State
	styles      Styles
}

type keySetConfig struct {
	align       int
	modifiers   []Modifier
	keyBindings []KeyBinding
	decorate    bool
	state       *commit.State
}

func newKeySet(ksc keySetConfig) []string {
	ks := keySet{
		align:       ksc.align,
		modifiers:   ksc.modifiers,
		keybindings: ksc.keyBindings,
		decorate:    ksc.decorate,
		styles:      defaultStyles(ksc.state.Theme),
		state:       ksc.state,
	}
	ks.keySetDimensions()
	ks.initKeySet()
	ks.fillKeySet()
	return ks.joinKeySet()
}

func (s *keySet) keySetDimensions() {
	s.height = s.countModifiers()

	width := 0
	for _, v := range s.modifiers {
		if v.Align != s.align {
			continue
		}
		i := s.countKeySet(v.Modifier)
		if i > width {
			width = i
		}
	}

	s.width = width
}

func (s *keySet) initKeySet() {
	col := make([][]string, s.height)

	for i := range col {
		row := make([]string, s.width*2)
		col[i] = row
	}

	s.target = col
}

func (s *keySet) fillKeySet() {
	i := 0
	for _, v := range s.modifiers {
		if v.Align != s.align {
			continue
		}

		j := 0
		for _, vv := range s.keybindings {
			if v.Modifier != vv.Modifier {
				continue
			}

			switch s.align {
			case AlignLeft:
				s.target[i][j] = vv.Key
				s.target[i][j+1] = vv.Label
			case AlignRight:
				s.target[i][j+1] = vv.Key
				s.target[i][j] = vv.Label
			}

			j += 2
		}

		i++
	}
}

func (s *keySet) joinKeySet() []string {
	var ss []string
	var offset int

	if len(s.target) == 0 {
		return []string{}
	}

	for i := 0; i < len(s.target[0]); i++ {
		var col []string
		for j := 0; j < len(s.target); j++ {
			col = append(col, s.target[j][i])
		}
		max := sliceMaxLen(col)

		str := s.joinColumn(col, max, offset)

		if str == "" {
			continue
		}

		m := s.styles.columnLeft.Render(str)
		if s.align == AlignRight {
			m = s.styles.columnRight.Render(str)
		}

		ss = append(ss, m)

		offset++
	}
	return ss
}

func (s *keySet) joinColumn(col []string, max int, offset int) string {
	var res []string

	remainder := 0
	lr := lipgloss.Right
	ll := lipgloss.Left

	if s.align == AlignRight {
		remainder = 1
		lr = lipgloss.Left
		ll = lipgloss.Right
	}

	for _, v := range col {
		switch offset%2 == remainder {
		case true:
			d := s.decorateKey(v, max, lr, s.decorate)
			res = append(res, d)
		case false:
			d := s.decorateLabel(v, max, ll, s.decorate)
			res = append(res, d)
		}
	}
	return strings.Join(res, "\n")
}

func (s *keySet) countModifiers() int {
	i := 0
	for _, m := range s.modifiers {
		if m.Align != s.align {
			continue
		}
		i++
	}
	return i
}

func (s *keySet) countKeySet(modifier int) int {
	i := 0
	for _, s := range s.keybindings {
		if s.Modifier != modifier {
			continue
		}
		i++
	}

	return i
}

func (s keySet) decorateKey(key string, max int, align lipgloss.Position, bracket bool) string {
	var k string
	padding := 0

	if key != "" {
		k = s.styles.key.Render(key)
		if bracket {
			padding = 2
			k = fmt.Sprintf("%v%v%v",
				s.styles.angleBracket.Render("<"),
				s.styles.key.Render(key),
				s.styles.angleBracket.Render(">"),
			)
		}
	}

	return lipgloss.NewStyle().Width(max + padding).Align(align).Render(k)
}

func (s keySet) decorateLabel(label string, max int, align lipgloss.Position, colour bool) string {
	var l string
	if label != "" {
		l = label
		if colour {
			l = s.styles.label.Render(label)
		}
	}

	return lipgloss.NewStyle().Width(max).Align(align).Render(l)
}

func sliceMaxLen(ss []string) int {
	var i int
	for _, v := range ss {
		if lipgloss.Width(v) > i {
			i = lipgloss.Width(v)
		}
	}

	return i
}
