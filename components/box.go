package components

import (
	"charm.land/lipgloss/v2"
	"github.com/Nick2k4L/BubbleIrc/theme"
	"github.com/charmbracelet/x/exp/charmtone"
)

// REFACTOR THIS -- THIS IS BAD AS WELL....

type Box struct {
	Width                int
	Height               int
	Padding              []int
	IsCenteredHorizontal bool
	IsSubBox             bool
	IsBlend              bool
	Theme                theme.Theme
	Content              string
	PopUpContent         string

	style lipgloss.Style
}

// Apply all the themeing props we need
func (b *Box) applyTheme() {
	// Start with a fresh base style
	b.style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

	if !b.IsSubBox {
		b.style = b.style.BorderForeground(lipgloss.Color(b.Theme.MainBorderColor))
	} else {
		b.style = b.style.BorderForeground(lipgloss.Color(b.Theme.SubBorderColor))
	}

	// this will need to change in the future....
	if b.IsBlend {
		b.style = b.style.BorderForegroundBlend(
			charmtone.Cherry,
			charmtone.Charple,
			charmtone.Guac,
			charmtone.Charple,
			charmtone.Sriracha,
		)
	}

	if b.IsCenteredHorizontal {
		b.style = b.style.AlignHorizontal(lipgloss.Center)
	}

	if b.Theme.BackgroundColor != "" {
		b.style = b.style.Background(lipgloss.Color(b.Theme.BackgroundColor)).
			BorderBackground(lipgloss.Color(b.Theme.BackgroundColor))
	}
}

// apply width & sizing
func (b *Box) applySizing() {
	if len(b.Padding) > 0 {
		b.style = b.style.Padding(b.Padding...)
	}
	if b.Width > 0 {
		b.style = b.style.Width(b.Width)
	}
	if b.Height > 0 {
		b.style = b.style.Height(b.Height)
	}
}

func (b *Box) Render() string {
	b.applyTheme()
	b.applySizing()
	return b.style.Render(b.Content)
}

func (b *Box) RenderPopUp(mainBox string) string {
	popup := lipgloss.NewLayer(b.Render())

	// Use the actual app UI as the base layer so the canvas matches the terminal size
	baseLayer := lipgloss.NewLayer(mainBox).AddLayers(
		popup.X(b.Width).Y(b.Height).Z(1),
	)

	comp := lipgloss.NewCompositor(baseLayer)
	return comp.Render()
}
