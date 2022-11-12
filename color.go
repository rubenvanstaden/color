package color

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func Colorize(color, s string) string {
	return color + s + Reset
}

func InRed(s string) string {
	return Colorize(Red, s)
}

func InGreen(s string) string {
	return Colorize(Green, s)
}

func InYellow(s string) string {
	return Colorize(Yellow, s)
}

func InBlue(s string) string {
	return Colorize(Blue, s)
}

func InPurple(s string) string {
	return Colorize(Purple, s)
}

func InCyan(s string) string {
	return Colorize(Cyan, s)
}

func InGray(s string) string {
	return Colorize(Gray, s)
}

func InWhite(s string) string {
	return Colorize(White, s)
}
