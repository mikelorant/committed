package setting

type Noop struct{}

func (n *Noop) Render(styles Styles) string {
	return ""
}

func (n *Noop) Focus() {}

func (n *Noop) Blur() {}

func (n *Noop) Type() Type {
	return TypeNoop
}
