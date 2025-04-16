package main

import (
	"regexp"
	"strconv"
)

func hexAndBinToDecimal(input string) string {
	for {
		converted := false
		// Преобразование bin
		binPattern := regexp.MustCompile(`\b(\w+)\s*\(\s*bin\s*\)`)
		input = binPattern.ReplaceAllStringFunc(input, func(match string) string {
			parts := binPattern.FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			binaryStr := parts[1]
			decimal, err := strconv.ParseInt(binaryStr, 2, 64)
			if err != nil {
				return match
			}
			converted = true
			return strconv.FormatInt(decimal, 10)
		})
		// Преобразование hex
		hexPattern := regexp.MustCompile(`\b(\w+)\s*\(\s*hex\s*\)`)
		input = hexPattern.ReplaceAllStringFunc(input, func(match string) string {
			parts := hexPattern.FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			hexStr := parts[1]
			decimal, err := strconv.ParseInt(hexStr, 16, 64)
			if err != nil || hexStr == "" {
				return match
			}
			converted = true
			return strconv.FormatInt(decimal, 10)
		})
		// Если преобразований больше нет, выходим из цикла
		if !converted {
			break
		}
	}

	// Удаляет остаточные опер
	leftoverPattern := regexp.MustCompile(`\(\s*(bin|hex)\s*\)`)
	leftovers := leftoverPattern.FindAllString(input, -1)
	if len(leftovers) > 0 {
		input = leftoverPattern.ReplaceAllString(input, "")
	}
	return input
}

// package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// )

// func hexAndBinToDecimal(input string) string {
// 	for {
// 		modified := false

// 		// replace hex conversions
// 		hexRegex := regexp.MustCompile(`(\w+)\s*\(hex\)`)
// 		input = hexRegex.ReplaceAllStringFunc(input, func(match string) string {
// 			parts := hexRegex.FindStringSubmatch(match)
// 			if len(parts) > 1 {
// 				hexVal, err := strconv.ParseInt(parts[1], 16, 64)
// 				if err == nil {
// 					modified = true
// 					return fmt.Sprintf("%d", hexVal)
// 				}
// 			}
// 			return match
// 		})

// 		// replace bin conversions
// 		binRegex := regexp.MustCompile(`(\w+)\s*\(bin\)`)
// 		input = binRegex.ReplaceAllStringFunc(input, func(match string) string {
// 			parts := binRegex.FindStringSubmatch(match)
// 			if len(parts) > 1 {
// 				binVal, err := strconv.ParseInt(parts[1], 2, 64)
// 				if err == nil {
// 					modified = true
// 					return fmt.Sprintf("%d", binVal)
// 				}
// 			}
// 			return match
// 		})

// 		// Если ничего не изменилось, выходим из цикла
// 		if !modified {
// 			break
// 		}
// 	}
// 	return input
// }
