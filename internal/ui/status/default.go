package status

func defaultModifiers() []Modifier {
	return []Modifier{
		{Modifier: NoModifier, Align: AlignRight},
		{Modifier: ShiftModifier, Label: "Shift", Align: AlignRight},
		{Modifier: AltModifier, Label: "Alt", Align: AlignLeft},
		{Modifier: ControlModifier, Label: "Ctrl", Align: AlignLeft},
	}
}

func defaultShortcuts() []Shortcut {
	return []Shortcut{
		{Modifier: ControlModifier, Key: "c", Label: "Cancel"},
		{Modifier: AltModifier, Key: "enter", Label: "Commit"},
		{Modifier: AltModifier, Key: "s", Label: "Sign-off"},
		{Modifier: AltModifier, Key: "/", Label: "Help"},
	}
}
