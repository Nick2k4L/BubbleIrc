package components

import "github.com/Nick2k4L/BubbleIrc/theme"

var (
	SubBox = Box{
		Width:    35,
		Theme:    theme.BaseTheme(),
		IsSubBox: true,
	}

	MainBox = Box{
		IsCenteredHorizontal: true,
		IsSubBox:             false,
		Theme:                theme.BaseTheme(),
	}

	PopUpBox0 = Box{
		IsCenteredHorizontal: true,
		IsSubBox:             false,
		Theme:                theme.BaseTheme(),
		IsBlend:              true,
	}
)
