package setting

import (
	"github.com/mikelorant/committed/internal/theme"
	"github.com/mikelorant/committed/internal/ui/colour"
)

type Styles struct{}

func defaultStyles(th theme.Theme) Styles {
	var s Styles

	_ = colour.New(th).OptionSetting()

	return s
}
