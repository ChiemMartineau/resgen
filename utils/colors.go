package utils

type Colour string

const (
	ColourBlack  Colour = "\x1b[0;30m"
	ColourRed    Colour = "\x1b[0;31m"
	ColourGreen  Colour = "\x1b[0;32m"
	ColourYellow Colour = "\x1b[0;33m"
	ColourBlue   Colour = "\x1b[0;34m"
	ColourPurple Colour = "\x1b[0;35m"
	ColourCyan   Colour = "\x1b[0;36m"
	ColourWhite  Colour = "\x1b[0;37m"
	ColourReset  Colour = "\x1b[0m"
)

const (
	ColourBlackBold  Colour = "\x1b[1;30m"
	ColourRedBold    Colour = "\x1b[1;31m"
	ColourGreenBold  Colour = "\x1b[1;32m"
	ColourYellowBold Colour = "\x1b[1;33m"
	ColourBlueBold   Colour = "\x1b[1;34m"
	ColourPurpleBold Colour = "\x1b[1;35m"
	ColourCyanBold   Colour = "\x1b[1;36m"
	ColourWhiteBold  Colour = "\x1b[1;37m"
)
