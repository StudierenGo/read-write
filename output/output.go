package output

import "github.com/fatih/color"

func PrintMessage(message any) {
	switch v := message.(type) {
	case string:
		color.Magenta(v)
	case int:
		color.Red("Error code %d", v)
	default:
		color.Yellow("Unsupported type")
	}
}
