package main

import (
	"strings"
)

func fixAtoAn(input string) string {
	silentHWords := map[string]bool{
		"honest":    true,
		"heir":      true,
		"honorific": true,
		"honor":     true,
		"herb":      true,
		"hotel":     true,
		"hour":      true,
		"homage":    true,
	}

	words := strings.Fields(input)
	for i := 0; i < len(words)-1; i++ {
		if len(words[i+1]) > 1 {
			if words[i] == "a" || words[i] == "A" {
				if silentHWords[words[i+1]] {
					if words[i] == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
					continue
				}
				firstLetter := []rune(words[i+1])[0]
				if strings.ContainsRune("aeiouAEIOU", firstLetter) {
					if words[i] == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			} else if words[i] == "an" || words[i] == "An" || words[i] == "AN" || words[i] == "aN" {
				if !silentHWords[words[i+1]] {
					firstLetter := []rune(words[i+1])[0]
					if strings.ContainsRune("bcdfgjklmnpqhrstvwxyzBCDFGJKLMNPQRSHTVWXYZ", firstLetter) {
						if words[i] == "an" {
							words[i] = "a"
						} else {
							words[i] = "A"
						}
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

// package main

// import (
// 	"strings"
// )

// func fixAtoAn(input string) string {
// 	words := strings.Fields(input)
// 	for i := 0; i < len(words)-1; i++ {
// 		if words[i] == "a" || words[i] == "A" {
// 			firstLetter := []rune(words[i+1])[0]
// 			if strings.ContainsRune("aeiouhAEIOUH", firstLetter) {
// 				if words[i] == "a" {
// 					words[i] = "an"
// 				} else {
// 					words[i] = "An"
// 				}
// 				words[i+1] = strings.ToLower(words[i+1])
// 			}
// 		} else if words[i] == "an" || words[i] == "An" || words[i] == "AN" || words[i] == "aN" {
// 			firstLetter := []rune(words[i+1])[0]
// 			if strings.ContainsRune("bcdfgjklmnpqrstvwxyazBCDFGJKL MNPQRSTVWXYAZ", firstLetter) {
// 				if words[i] == "an" {
// 					words[i] = "a"
// 				} else {
// 					words[i] = "A"
// 				}
// 				words[i+1] = strings.ToLower(words[i+1])
// 			}
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
