package core

import (
	"fmt"
	http "net/http"
)

// colors
const (
	green   = "\033[32m"
	white   = "\033[37m"
	yellow  = "\033[33m"
	red     = "\033[31m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	reset   = "\033[0m"
)

var (
	register = Colorize(green, "[Register]")
	starting = Colorize(cyan, "[Starting...]")
)

func Colorize(color string, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, reset)
}

func GetStatusString(status HTTPMethods) string {
	switch status {
	case http.MethodGet:
		return Colorize(blue, fmt.Sprintf("[%s]", string(http.MethodGet)))
	case http.MethodPost:
		return Colorize(cyan, fmt.Sprintf("[%s]", string(http.MethodPost)))
	case http.MethodDelete:
		return Colorize(red, fmt.Sprintf("[%s]", string(http.MethodDelete)))
	case http.MethodPut:
		return Colorize(yellow, fmt.Sprintf("[%s]", string(http.MethodPut)))
	case http.MethodPatch:
		return Colorize(green, fmt.Sprintf("[%s]", string(http.MethodPatch)))
	default:
		return string(status)
	}
}
