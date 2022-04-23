package cli

import (
	"time"

	"github.com/fatih/color"
)

// Message is used to print a message to the command line
func message(level string, message string) {
	switch level {
	case "info":
		color.Cyan("[*]" + message)
	case "note":
		color.Yellow("[-]" + message)
	case "warn":
		color.Red("[!]" + message)
	case "debug":
		color.Red("[DEBUG]" + message)
	case "success":
		color.Green("[+]" + message)
	default:
		color.Red("[_-_]Invalid message level: " + message)
	}
}

func ListCmdResult() {
	beat := time.Tick(2 * time.Second)
	for range beat {
		DisplayCmdResult()
	}
}