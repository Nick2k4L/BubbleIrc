package components

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

// CreatePopUpBox creates out popbox, so we need to take in
// 1. whatever the screen is currently displaying
// 2. whatever we want in the content, for example
func CreatePopUpBox(mainAppContent string, popupContent string) string {
	popup := lipgloss.NewLayer(box(popupContent))

	// Use the actual app UI as the base layer so the canvas matches the terminal size
	baseLayer := lipgloss.NewLayer(mainAppContent).AddLayers(
		popup.X(55).Y(10).Z(1),
	)

	comp := lipgloss.NewCompositor(baseLayer)
	return comp.Render()
}

func box(content string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForegroundBlend(
			charmtone.Cherry,
			charmtone.Charple,
			charmtone.Guac,
			charmtone.Charple,
			charmtone.Sriracha,
		).Background(color.White).BorderBackground(color.White).
		Height(20).
		Width(60).
		AlignHorizontal(lipgloss.Center).
		Render(content)
}
