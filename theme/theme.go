package theme

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"
)

// eventually want to take advantage of toml configs to allow users to create their own themes
// and use predefined themes within a .toml file
// we could also expand to language translations in the future

type Theme struct {
	BackgroundColor string
	MainBorderColor string
	SubBorderColor  string
	ForeGroundColor string
	FocusedStyle    string
	BlurredStyle    string
	UnFocusedStyle  string
	InputChar       string
	ForeGroundBlend []string // if we want ForeGroundBlend --- need to modify this
}

type InputModel struct{}

// Basic theme creation
func CreateTheme(backgroundColor, mainBorderColor,
	subBorderColor, foreGroundColor, focusedStyle, unFocusedStyle string, inputChar string, foreGroundBlend ...string,
) Theme {
	return Theme{
		BackgroundColor: backgroundColor,
		MainBorderColor: mainBorderColor,
		SubBorderColor:  subBorderColor,
		ForeGroundColor: foreGroundColor,
		FocusedStyle:    focusedStyle,
		UnFocusedStyle:  unFocusedStyle,
		InputChar:       inputChar,
		ForeGroundBlend: foreGroundBlend,
	}
}

func BaseTheme() Theme {
	return CreateTheme(
		"#1E1E1E", // backgroundColor: Dark charcoal
		"#00ADD8", // mainBorderColor: Cyan
		"#444444", // subBorderColor: Muted gray
		"#FFFFFF", // foreGroundColor: Pure white
		"#FF007F", // focusedStyle: Bright pink
		"#777777", // unFocusedStyle: Dim gray
		"* ",
		"#CCCCCC", // foreGroundBlend: Light gray
		"#999999", // foreGroundBlend: Medium gray
	)
}

func (t Theme) CreateInputBoxModelTheme(textInput textinput.Model, s textinput.Styles) (textinput.Model, textinput.Styles) {
	focusedStyle := lipgloss.NewStyle().Foreground((lipgloss.Color(t.FocusedStyle))).Background(lipgloss.Color(t.BackgroundColor))
	blurredStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(t.UnFocusedStyle)).Background(lipgloss.Color(t.BackgroundColor))
	// Prompt color i.e background 'char' of character we end up choosing
	s.Focused.Prompt = focusedStyle
	s.Blurred.Prompt = blurredStyle

	// Text colors
	s.Focused.Text = focusedStyle
	s.Blurred.Text = blurredStyle

	// Color for blinking cursor
	s.Cursor.Color = lipgloss.Color(t.UnFocusedStyle)

	// background color for placeholder
	s.Focused.Placeholder = blurredStyle
	s.Blurred.Placeholder = blurredStyle

	textInput.Prompt = t.InputChar // Can use this to create custom prompt '>' indicators later...
	return textInput, s
}
