package utils

import (
	"fmt"
	"strings"
)

// FindVowels will check if the character is one of a, e, i, o, u
func FindVowels(s string) int {
	result := 0
	for _, v := range s {
		if strings.ContainsAny("aeiou", string(v)) {
			result++
		}
	}
	return result
}

// FindDigits will check only for digits from 0 to 9
func FindDigits(s string) int {
	result := 0
	for _, v := range s {
		if strings.ContainsAny("1234567890", string(v)) {
			result++
		}
	}
	return result
}

// FindSymbols will count common symbols appearing on the number keys
func FindSymbols(s string) int {
	result := 0
	for _, v := range s {
		if strings.ContainsAny("!@€£#$%^&*()", string(v)) {
			result++
		}
	}
	return result
}

// FindWords will count the number of words without surrounding spaces
func FindWords(s string) int {
	return len(strings.Fields(s))
}

// PrintReport will display a summary of results
func PrintReport(vowels, digits, symbols, words int) {
	fmt.Println("Analysis Summary")
	fmt.Println("-------- -------")
	fmt.Println("vowels  :", vowels)
	fmt.Println("digits  :", digits)
	fmt.Println("symbols :", symbols)
	fmt.Println("words   :", words)
}
