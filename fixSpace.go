package main

import (
	"regexp"
	"strings"
)

func fixSpace(input string) string {
	input = regexp.MustCompile(`((\()[a-z1-9])`).ReplaceAllString(input, " $1")
	input = regexp.MustCompile(`([a-z1-9](\)))`).ReplaceAllString(input, "$1 ")

	return strings.TrimSpace(input)
}

func processString(input string) string {
	for {
		modified := false
		newInput := textModifyCase(input)
		if newInput != input {
			input = newInput
			modified = true
		}
		newInput = hexAndBinToDecimal(input)
		if newInput != input {
			input = newInput
			modified = true
		}
		if !modified {
			break
		}
	}

	return input
}
