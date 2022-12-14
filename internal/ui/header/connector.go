package header

import (
	"fmt"
	"strings"
)

type connectorStyle string

const (
	connecterTopRight    connectorStyle = "└"
	connectopTopLeft     connectorStyle = "┘"
	connectorBottomLeft  connectorStyle = "┐"
	connectorBottomRight connectorStyle = "┌"
	connectorLeftTop     connectorStyle = "└"
	connectorLeftBottom  connectorStyle = "┌"
	connectorRightBottom connectorStyle = "┐"
	connectorRightTop    connectorStyle = "┘"
	connectorHorizonal   connectorStyle = "─"
	connectorVertical    connectorStyle = "│"
)

func connector(start, axes, end connectorStyle, len int) string {
	return fmt.Sprintf("%v%v%v", start, strings.Repeat(string(axes), len), end)
}
