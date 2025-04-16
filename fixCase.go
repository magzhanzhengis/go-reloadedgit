package main

import (
	"regexp"
	"strconv"
	"strings"
)

func textModifyCase(input string) string {
	re := regexp.MustCompile(`([^\(\)\n]*?)\s*?\(\s*?(\s*up|low|cap\s*)(?:,\s*(-?\d+))?\s*\)`)

	match := re.FindStringSubmatchIndex(input)
	if match == nil {
		return strings.TrimSpace(input)
	}

	textStart := match[2]
	textEnd := match[3]
	opStart := match[4]
	opEnd := match[5]
	countStart := match[6]
	countEnd := match[7]

	text := input[textStart:textEnd]
	operation := strings.TrimSpace(input[opStart:opEnd])
	count := 1

	if countStart != -1 {
		countStr := input[countStart:countEnd]
		if n, err := strconv.Atoi(countStr); err == nil {
			if n < 0 {
				// Оставляем текст без изменений и удаляем оператор
				newInput := input[:match[0]] + text + input[match[1]:]
				return textModifyCase(newInput)
			} else {
				count = n
			}
		}
	}

	words := strings.Split(text, " ")
	if count > len(words) {
		count = len(words)
	}
	startWord := len(words) - count
	if startWord < 0 {
		startWord = 0
	}

	for i := startWord; i < len(words); i++ {
		switch operation {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = Capitalize(words[i])
		}
	}

	modifiedText := strings.Join(words, " ")
	newInput := input[:match[0]] + modifiedText + input[match[1]:]

	return textModifyCase(newInput)
}

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	if len(word) == 1 {
		return strings.ToUpper(string(word[0]))
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// package main

// import (
// 	"regexp"
// 	"strconv"
// 	"strings"
// )

// func textModifyCase(input string) string {
// 	re := regexp.MustCompile(`(\w+(?:\s+\w+)*?)\s*\((up|low|cap)(?:,\s*(\d+))?\)`)
// 	return re.ReplaceAllStringFunc(input, func(match string) string {
// 		parts := re.FindStringSubmatch(match)
// 		if len(parts) < 4 {
// 			return match
// 		}

// 		text := strings.Fields(parts[1]) // Use Fields to handle multiple whitespaces
// 		operation := parts[2]            // type operation: up, low, cap

// 		// Parse count with default handling
// 		count := 1
// 		if parts[3] != "" {
// 			if n, err := strconv.Atoi(parts[3]); err == nil {
// 				count = n
// 			}
// 		}

// 		// Limit count to text length
// 		start := len(text) - count
// 		if start < 0 {
// 			start = 0
// 		}
// 		// Apply operation to specified number of words
// 		for i := start; i < len(text); i++ {
// 			switch operation {
// 			case "up":
// 				text[i] = strings.ToUpper(text[i])
// 			case "low":
// 				text[i] = strings.ToLower(text[i])
// 			case "cap":
// 				text[i] = capitalize(text[i])
// 			}
// 		}

// 		return strings.Join(text, " ")
// 	})
// }

// func capitalize(word string) string {
// 	if len(word) == 0 {
// 		return word
// 	}
// 	if len(word) == 1 {
// 		return strings.ToUpper(word)
// 	}
// 	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
// }
