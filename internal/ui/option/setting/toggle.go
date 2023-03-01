package setting

import (
	"fmt"
	"strings"
)

type Toggle struct {
	Title  string
	Enable bool

	focus bool
}

func (t *Toggle) Render(styles Styles) string {
	var str []string

	switch t.focus {
	case true:
		str = append(str, styles.settingTitleSelected.Render(t.Title))
	default:
		str = append(str, styles.settingTitle.Render(t.Title))
	}

	switch t.Enable {
	case true:
		v := styles.settingSelected.Render("Enable")
		str = append(str, fmt.Sprintf("%v %v", styles.settingSquareFilled, v))
	default:
		v := styles.setting.Render("Enable")
		str = append(str, fmt.Sprintf("%v %v", styles.settingSquareEmpty, v))
	}

	return strings.Join(str, "\n")
}

func (t *Toggle) Focus() {
	t.focus = true
}

func (t *Toggle) Blur() {
	t.focus = false
}

func (t *Toggle) Value() bool {
	return t.Enable
}

func (t *Toggle) Toggle() {
	t.Enable = !t.Enable
}

func (t *Toggle) Type() Type {
	return TypeToggle
}

func ToToggle(p Paner) *Toggle {
	return p.(*Toggle)
}
