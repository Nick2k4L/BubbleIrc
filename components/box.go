package components

import (
	"charm.land/lipgloss/v2"
	"github.com/Nick2k4L/BubbleIrc/theme"
	"github.com/charmbracelet/x/exp/charmtone"
)

// Use styles
func MainBoxStyle(th theme.Theme) lipgloss.Style {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(th.MainBorderColor))
	if th.BackgroundColor != "" {
		style = style.Background(lipgloss.Color(th.BackgroundColor)).
			BorderBackground(lipgloss.Color(th.BackgroundColor))
	}

	return style
}

func SubBoxStyle(th theme.Theme) lipgloss.Style {
	return MainBoxStyle(th).
		BorderForeground(lipgloss.Color(th.SubBorderColor)).
		Width(35)
}

func PopUpBoxStyle(th theme.Theme) lipgloss.Style {
	return MainBoxStyle(th).
		BorderForegroundBlend(
			charmtone.Cherry,
			charmtone.Charple,
			charmtone.Guac,
			charmtone.Charple,
			charmtone.Sriracha,
		)
}

func RenderPopUp(popUpBox, mainBox string, width, height int) string {
	popup := lipgloss.NewLayer(popUpBox)

	// Use the actual app UI as the base layer so the canvas matches the terminal size
	baseLayer := lipgloss.NewLayer(mainBox).AddLayers(
		popup.X(width).Y(height).Z(1),
	)

	comp := lipgloss.NewCompositor(baseLayer)
	return comp.Render()
}
