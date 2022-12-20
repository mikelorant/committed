package status

func defaultModifiers() []Modifier {
	return []Modifier{
		{Modifier: noModifier, Align: alignRight},
		{Modifier: shiftModifier, Label: "Shift", Align: alignRight},
		{Modifier: altModifier, Label: "Alt", Align: alignLeft},
		{Modifier: controlModifier, Label: "Ctrl", Align: alignLeft},
	}
}

func defaultShortcuts() []Shortcut {
	return []Shortcut{
		{Modifier: controlModifier, Key: "c", Label: "Cancel"},
		{Modifier: altModifier, Key: "enter", Label: "Commit"},
		{Modifier: altModifier, Key: "s", Label: "Sign-off"},
		{Modifier: altModifier, Key: "/", Label: "Help"},
	}
}
