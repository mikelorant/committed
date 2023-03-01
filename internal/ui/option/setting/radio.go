package setting

import (
	"fmt"
	"strings"
)

type Radio struct {
	Title  string
	Values []string
	Index  int

	focus bool
}

func (r *Radio) Render(styles Styles) string {
	var str []string

	switch r.focus {
	case true:
		str = append(str, styles.settingTitleSelected.Render(r.Title))
	default:
		str = append(str, styles.settingTitle.Render(r.Title))
	}

	for idx, val := range r.Values {
		if idx == r.Index {
			v := styles.settingSelected.Render(val)
			str = append(str, fmt.Sprintf("%v %v", styles.settingDotFilled, v))

			continue
		}

		v := styles.setting.Render(val)
		str = append(str, fmt.Sprintf("%v %v", styles.settingDotEmpty, v))
	}

	return strings.Join(str, "\n")
}

func (r *Radio) Focus() {
	r.focus = true
}

func (r *Radio) Blur() {
	r.focus = false
}

func (r *Radio) Value() string {
	return r.Values[r.Index]
}

func (r *Radio) Next() {
	if r.Index >= len(r.Values)-1 {
		r.Index = 0
		return
	}

	r.Index++
}

func (r *Radio) Previous() {
	if r.Index <= 0 {
		r.Index = len(r.Values) - 1
		return
	}

	r.Index--
}

func (r *Radio) Select(i int) {
	r.Index = i
}

func (r *Radio) Type() Type {
	return TypeRadio
}

func ToRadio(p Paner) *Radio {
	r, _ := p.(*Radio)

	return r
}
