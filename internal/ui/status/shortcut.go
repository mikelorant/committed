package status

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type shortcuts struct {
	shortcuts []Shortcut
	modifiers []Modifier
	styles    Styles
	view      string
}

type shortcutSet struct {
	align     int
	shortcuts []Shortcut
	modifiers []Modifier
	target    [][]string
	decorate  bool
	height    int
	width     int
	styles    Styles
}

type Shortcut struct {
	Modifier int
	Key      string
	Label    string
}

type Modifier struct {
	Modifier int
	Label    string
	Align    int
}

const (
	noModifier = iota
	altModifier
	controlModifier
	shiftModifier

	alignLeft = iota
	alignRight
)

func newShortcuts() shortcuts {
	s := shortcuts{
		modifiers: defaultModifiers(),
		shortcuts: defaultShortcuts(),
		styles:    defaultStyles(),
	}
	s.view = s.render()

	return s
}

func newShortcutSet(a int, ms []Modifier, ss []Shortcut, d bool) []string {
	scs := shortcutSet{
		align:     a,
		modifiers: ms,
		shortcuts: ss,
		decorate:  d,
		styles:    defaultStyles(),
	}
	scs.shortcutDimensions()
	scs.initShortcuts()
	scs.fillShortcuts()
	return scs.joinShortcuts()
}

func (s shortcuts) render() string {
	l := newShortcutSet(alignLeft, s.modifiers, s.shortcuts, true)
	r := newShortcutSet(alignRight, s.modifiers, s.shortcuts, true)
	ml := newShortcutSet(alignLeft, s.modifiers, modifierToShortcut(alignLeft, s.modifiers), false)
	mr := newShortcutSet(alignRight, s.modifiers, modifierToShortcut(alignRight, s.modifiers), false)

	var left, right []string

	if len(l) != 0 {
		left = append(left, ml...)
	}
	left = append(left, l...)

	right = append(right, r...)
	if len(r) != 0 {
		right = append(right, mr...)
	}

	hleft := lipgloss.JoinHorizontal(lipgloss.Top, left...)
	hright := lipgloss.JoinHorizontal(lipgloss.Top, right...)

	bleft := s.styles.shortcutBlockLeft.Render(hleft)
	bright := s.styles.shortcutBlockRight.Render(hright)

	block := lipgloss.JoinHorizontal(lipgloss.Top, bleft, bright)

	return s.styles.shortcutBoundary.Render(block)
}

func (s *shortcutSet) shortcutDimensions() {
	s.height = s.countModifiers()

	width := 0
	for _, v := range s.modifiers {
		if v.Align != s.align {
			continue
		}
		i := s.countShortcuts(v.Modifier)
		if i > width {
			width = i
		}
	}

	s.width = width
}

func (s *shortcutSet) initShortcuts() {
	col := make([][]string, s.height)

	for i := range col {
		row := make([]string, s.width*2)
		col[i] = row
	}

	s.target = col
}

func (s *shortcutSet) fillShortcuts() {
	i := 0
	for _, v := range s.modifiers {
		if v.Align != s.align {
			continue
		}

		j := 0
		for _, vv := range s.shortcuts {
			if v.Modifier != vv.Modifier {
				continue
			}

			switch s.align {
			case alignLeft:
				s.target[i][j] = vv.Key
				s.target[i][j+1] = vv.Label
			case alignRight:
				s.target[i][j+1] = vv.Key
				s.target[i][j] = vv.Label
			}

			j += 2
		}

		i++
	}
}

func (s *shortcutSet) joinShortcuts() []string {
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
		len := sliceMaxLen(col)

		str := s.joinColumn(col, len, offset)

		m := s.styles.shortcutColumnRight.Render(str)
		if s.align == alignRight {
			m = s.styles.shortcutColumnLeft.Render(str)
		}

		ss = append(ss, m)

		offset++
	}
	return ss
}

func (s *shortcutSet) joinColumn(col []string, len int, offset int) string {
	var res []string

	remainder := 0
	lr := lipgloss.Right
	ll := lipgloss.Left

	if s.align == alignRight {
		remainder = 1
		lr = lipgloss.Left
		ll = lipgloss.Right
	}

	for _, v := range col {
		switch offset%2 == remainder {
		case true:
			d := s.decorateKey(v, len, lr, s.decorate)
			res = append(res, d)
		case false:
			d := s.decorateLabel(v, len, ll, s.decorate)
			res = append(res, d)
		}
	}
	return strings.Join(res, "\n")
}

func (s *shortcutSet) countModifiers() int {
	i := 0
	for _, m := range s.modifiers {
		if m.Align != s.align {
			continue
		}
		i++
	}
	return i
}

func (s *shortcutSet) countShortcuts(modifier int) int {
	i := 0
	for _, s := range s.shortcuts {
		if s.Modifier != modifier {
			continue
		}
		i++
	}

	return i
}

func modifierToShortcut(a int, ms []Modifier) []Shortcut {
	var ss []Shortcut
	var label string

	for _, v := range ms {
		if v.Align != a {
			continue
		}

		if v.Label != "" {
			label = "+"
		}

		scs := Shortcut{
			Modifier: v.Modifier,
			Key:      v.Label,
			Label:    label,
		}

		ss = append(ss, scs)
	}

	return ss
}

func (s shortcutSet) decorateKey(key string, len int, align lipgloss.Position, bracket bool) string {
	var k string
	padding := 0

	if key != "" {
		k = s.styles.shortcutKey.Render(key)
		if bracket {
			padding = 2
			k = fmt.Sprintf("<%s>", s.styles.shortcutKey.Render(key))
		}
	}

	return lipgloss.NewStyle().Width(len + padding).Align(align).Render(k)
}

func (s shortcutSet) decorateLabel(label string, len int, align lipgloss.Position, colour bool) string {
	var l string
	if label != "" {
		l = label
		if colour {
			l = s.styles.shortcutLabel.Render(label)
		}
	}

	return lipgloss.NewStyle().Width(len).Align(align).Render(l)
}

func sliceMaxLen(ss []string) int {
	var i int
	for _, v := range ss {
		if len(v) > i {
			i = len(v)
		}
	}

	return i
}
