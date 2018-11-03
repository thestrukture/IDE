package methods

import "strings"

//
func ParseLog(args ...interface{}) string {
	cline := args[0]

	calls := strings.Split(cline.(string), ":")
	actionText := ""
	if calls[0] == "service" {
		actionText = "The line is located in  Web service ( " + calls[1] + ") at line: " + calls[2]
	} else if calls[0] == "init" {
		actionText = "The line is located in your package Init func at line: " + calls[1]
	} else if calls[0] == "main" {
		actionText = "The line is located in your package Main func at line: " + calls[1]
	} else if calls[0] == "structs" {
		actionText = "The line is located in your package Interfaces at line: " + calls[1]
	} else if calls[0] == "meth" {
		actionText = "The line is located in your package template pipelines at line: " + calls[1]
	}
	return actionText

}
