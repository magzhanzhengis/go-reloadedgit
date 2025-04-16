package main

import (
	"regexp"
	"strings"
)

func fixPunctuations(input string) string {
	// Убираем пробелы вокруг знаков пунктуации
	re := regexp.MustCompile(`\s*([.,!?:;])\s*`)
	input = re.ReplaceAllString(input, "$1")
	// Добавляем пробел после знаков пунктуации, если следом идет буква или цифра
	re = regexp.MustCompile(`([.,!?:;])([a-zA-Z0-9-])`)
	input = re.ReplaceAllString(input, "$1 $2")
	// Убираем двойные пробелы
	input = strings.Join(strings.Fields(input), " ")
	// Убираем начальные и конечные пробелы
	return strings.TrimSpace(input)
}

// Функция для обработки двойных кавычек
func fixDoubleQuotes(input string) string {
	// Убираем пробелы внутри двойных кавычек
	re := regexp.MustCompile(`"\s*(.*?)\s*"`)
	input = re.ReplaceAllString(input, `"$1"`)
	// Убираем пробелы между кавычкой и прилегающим словом
	re = regexp.MustCompile(`(["])\s+([\'\w])`)
	input = re.ReplaceAllString(input, `$1$2`)
	re = regexp.MustCompile(`([\'\w])\s+(["])`)
	input = re.ReplaceAllString(input, `$1$2`)
	// Подсчитываем общее количество кавычек
	quoteCount := strings.Count(input, `"`)
	var result []rune
	inQuote := false
	for i := 0; i < len(input); i++ {
		currentChar := rune(input[i])
		if currentChar == '"' {
			if quoteCount%2 != 0 && strings.Count(string(result), `"`) == quoteCount-1 {
				result = append(result, currentChar, ' ')
				continue
			}
			if !inQuote {
				inQuote = true
				if i > 0 && input[i-1] != ' ' && input[i-1] != '"' && input[i-1] != '\'' {
					result = append(result, ' ')
				}
				result = append(result, currentChar)
			} else {
				inQuote = false
				result = append(result, currentChar)
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
			}
		} else {
			result = append(result, currentChar)
		}
	}
	return strings.TrimSpace(string(result))
}

func fixSingleQuotes(input string) string {
	// Убираем пробелы внутри одинарных кавычек
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	input = re.ReplaceAllString(input, "'$1'")
	// Убираем пробелы между кавычкой и прилегающим словом
	re = regexp.MustCompile(`(['])\s+([\'\w])`)
	input = re.ReplaceAllString(input, `$1$2`)
	re = regexp.MustCompile(`([\'\w])\s+(['])`)
	input = re.ReplaceAllString(input, `$1$2`)
	validSuffixes := map[string]bool{
		"t":  true,
		"ll": true,
		"ve": true,
		"m":  true,
		"s":  true,
		"d":  true,
		"re": true,
	}
	quoteCount := strings.Count(input, "'")
	var result []rune
	inQuote := false
	for i := 0; i < len(input); i++ {
		currentChar := rune(input[i])
		if currentChar == '\'' {
			if i > 0 && i+1 < len(input) {
				prevWordEnd := i - 1
				for prevWordEnd >= 0 && isLetter(rune(input[prevWordEnd])) {
					prevWordEnd--
				}
				prevWord := input[prevWordEnd+1 : i]
				nextWordStart := i + 1
				for nextWordStart < len(input) && isLetter(rune(input[nextWordStart])) {
					nextWordStart++
				}
				nextSuffix := input[i+1 : nextWordStart]
				if isWord(prevWord) && validSuffixes[nextSuffix] {
					result = append(result, currentChar)
					continue
				}
			}
			if quoteCount == 1 {
				result = append(result, currentChar)
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
				continue
			}
			if quoteCount%2 != 0 && strings.Count(string(result), `'`) == quoteCount-1 && !inQuote {
				result = append(result, currentChar)
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
				continue
			}
			if !inQuote {
				inQuote = true
				if i > 0 && input[i-1] != ' ' && input[i-1] != '\'' && input[i-1] != '"' {
					result = append(result, ' ')
				}
				result = append(result, currentChar)
			} else {
				inQuote = false
				result = append(result, currentChar)

				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
			}
		} else {
			result = append(result, currentChar)
		}
	}
	return strings.TrimSpace(string(result))
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isWord(s string) bool {
	for _, ch := range s {
		if !isLetter(ch) {
			return false
		}
	}
	return len(s) > 0
}

// package main

// import (
// 	"regexp"
// 	"strings"
// )

// func fixPunctuations(input string) string {
// 	// Обрабатываем пробелы вокруг всех указанных знаков пунктуации
// 	re := regexp.MustCompile(`\s*([.,!?:;'"])\s*`)
// 	input = re.ReplaceAllString(input, "$1")

// 	// Добавляем пробел между словами и кавычками
// 	re = regexp.MustCompile(`(")(\w+)(")(\s*)(")(\w+)(")`)
// 	input = re.ReplaceAllString(input, "$1$2$3 $5$6$7")

// 	// Добавляем пробел после .,!?;
// 	re = regexp.MustCompile(`([.,!?;])([a-zA-Z0-9])`)
// 	input = re.ReplaceAllString(input, "$1 $2")

// 	// Убираем двойные кавычки ' и ' с пробелами внутри
// 	re = regexp.MustCompile(`'\s*([^']*)\s*'`)
// 	input = re.ReplaceAllString(input, "'$1'")

// 	// Добавляем пробел после двоеточия, если за ним идет слово
// 	re = regexp.MustCompile(`(:)([.,!?:;"'a-zA-Z0-9])`)
// 	input = re.ReplaceAllString(input, "$1 $2")

// 	re = regexp.MustCompile(`'(\w+)'`)
// 	input = re.ReplaceAllString(input, " '$1' ")

// 	// Убираем двойные пробелы
// 	input = strings.ReplaceAll(input, "  ", " ")

// 	re = regexp.MustCompile(`(["'w+'"])(["'w+'"])`)
// 	input = re.ReplaceAllString(input, "$1 $2")

// 	re = regexp.MustCompile(`(["'])([^"']+)([,"'])([a-zA-Z])`)
// 	input = re.ReplaceAllString(input, "$1$2$3 $4")

// 	re = regexp.MustCompile(`(['])([t])`)
// 	input = re.ReplaceAllString(input, "$1$2")

// 	// Обработка апострофа с "t"
// 	re = regexp.MustCompile(`\s*'\s*t\s*`)
// 	input = re.ReplaceAllString(input, "'t ")

// 	// Убираем начальные и конечные пробелы
// 	return strings.TrimSpace(input)
// }
