package log4go
import "regexp";

const (
	//BLACK = "\x1B[30m"
	RED = "\x1B[31m"
	GREEN = "\x1B[32m"
	YELLOW = "\x1B[33m"
	BLUE = "\x1B[34m"
	MAGENTA = "\x1B[35m"
	CYAN = "\x1B[36m"
	WHITE = "\x1B[37m"
DEFAULT = "\x1B[39m"
	BOLD = "\x1B[1m"
	RESET = "\x1B[0m"
)

var reset, bold, blue, red, green, yellow, magenta, cyan, white, defaultcolor *regexp.Regexp

func init() {
	reset, _ = regexp.Compile("%R")
	bold, _ = regexp.Compile("%e")
	blue, _ = regexp.Compile("%b")
	red, _ = regexp.Compile("%r")
	green, _ = regexp.Compile("%g")
	yellow, _ = regexp.Compile("%y")
	magenta, _ = regexp.Compile("%m")
	cyan, _ = regexp.Compile("%c")
	white, _ = regexp.Compile("%w")
	defaultcolor, _ = regexp.Compile("%d")
}

func Colorize(str string) (answer string) {
	answer = reset.ReplaceAllString(str, RESET)
	answer = bold.ReplaceAllString(answer, BOLD)
	answer = blue.ReplaceAllString(answer, BLUE)
	answer = red.ReplaceAllString(answer, RED)
	answer = green.ReplaceAllString(answer, GREEN)
	answer = yellow.ReplaceAllString(answer, YELLOW)
	answer = magenta.ReplaceAllString(answer, MAGENTA)
	answer = cyan.ReplaceAllString(answer, CYAN)
	answer = white.ReplaceAllString(answer, WHITE)
	answer = defaultcolor.ReplaceAllString(answer, DEFAULT)
	return
}