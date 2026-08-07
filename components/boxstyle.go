package components

import "charm.land/lipgloss/v2"

var (
	Pink   = "#fd4be8"
	Teal   = "#4bfde5"
	Orange = "#fd794b"
)

// BoxOptions. This allows to create custom boxes / borders as we go on
// With things such as themes in mind, this will come in hand
type BoxOptions struct {
	Color   string
	Padding []int
	Width   int
	Height  int
	Title   string
}

// CreateNewRoundedBoxStyle will create a new rounded box
func CreateNewRoundedBoxStyle(boxOptions BoxOptions) lipgloss.Style {
	// rounded box
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(boxOptions.Color))

	// we check if any of these options are set, if they are not move on
	if len(boxOptions.Padding) > 0 {
		boxStyle.Padding(boxOptions.Padding...)
	}

	if boxOptions.Width > 0 {
		boxStyle.Width(boxOptions.Width)
	}

	if boxOptions.Height > 0 {
		boxStyle.Height(boxOptions.Height)
	}

	return boxStyle
}
