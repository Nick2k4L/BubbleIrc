package theme

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
	ForeGroundBlend []string // if we want ForeGroundBlend --- need to modify this
}

// Basic theme creation
func CreateTheme(backgroundColor, mainBorderColor,
	subBorderColor, foreGroundColor, focusedStyle, unFocusedStyle string, foreGroundBlend ...string,
) Theme {
	return Theme{
		BackgroundColor: backgroundColor,
		MainBorderColor: mainBorderColor,
		SubBorderColor:  subBorderColor,
		ForeGroundColor: foreGroundColor,
		FocusedStyle:    focusedStyle,
		UnFocusedStyle:  unFocusedStyle,
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
		"#CCCCCC", // foreGroundBlend: Light gray
		"#999999", // foreGroundBlend: Medium gray
	)
}
