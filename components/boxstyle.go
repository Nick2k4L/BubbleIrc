package components

import (
	"charm.land/lipgloss/v2"
	"github.com/Nick2k4L/BubbleIrc/theme"
)

var (
	Pink            = "#fd4be8"
	Teal            = "#4bfde5"
	Orange          = "#fd794b"
	BackgroundColor = "#ffffff"
)

// BoxOptions. This allows to create custom boxes / borders as we go on
// With things such as themes in mind, this will come in hand
type BoxOptions struct {
	// Color                string
	Padding []int
	Width   int
	Height  int
	// Title                string
	// BackgroundColor      color.Color
	IsCenteredHorizontal bool
	IsSubBox             bool
}

// CreateNewRoundedBoxStyle will create a new rounded box -- boxOptions should eventually consume theme
func CreateNewRoundedBoxStyle(content string, boxOptions BoxOptions, theme theme.Theme) string {
	// rounded box
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()) // we want this to be a rounded box -- keep this as is

	if !boxOptions.IsSubBox {
		boxStyle = boxStyle.BorderForeground(lipgloss.Color(theme.MainBorderColor))
	} else {
		boxStyle = boxStyle.BorderForeground(lipgloss.Color(theme.SubBorderColor))
	}

	if boxOptions.IsCenteredHorizontal {
		boxStyle = boxStyle.AlignHorizontal(lipgloss.Center)
	}

	// we check if any of these options are set, if they are not move on
	if len(boxOptions.Padding) > 0 {
		boxStyle = boxStyle.Padding(boxOptions.Padding...)
	}

	if boxOptions.Width > 0 {
		boxStyle = boxStyle.Width(boxOptions.Width)
	}

	if boxOptions.Height > 0 {
		boxStyle = boxStyle.Height(boxOptions.Height)
	}

	if theme.BackgroundColor != "" {
		boxStyle = boxStyle.Background(lipgloss.Color(theme.BackgroundColor)).BorderBackground(lipgloss.Color(theme.BackgroundColor))
	}

	return boxStyle.Render(content)
}
