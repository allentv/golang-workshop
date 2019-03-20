package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputStr := flag.String("text", "", "Text block to parse")
	flag.Parse()

	cleanedStr := cleanupInput(*inputStr)

	vowels := findVowels(cleanedStr)
	digits := findDigits(cleanedStr)
	symbols := findSymbols(cleanedStr)
	words := findWords(cleanedStr)

	printReport(vowels, digits, symbols, words)
}

// cleanupInput runs all the validation and bleaching to be applied on the input string
func cleanupInput(inputStr string) string {
	if inputStr == "" {
		log.Fatal("Input string not provided!")
		os.Exit(1)
	}
	return strings.ToLower(inputStr)
}

// findVowels will check if the character is one of a, e, i, o, u
func findVowels(s string) int {
	result := 0
	for _, v := range s {
		if strings.ContainsAny("aeiou", string(v)) {
			result++
		}
	}
	return result
}

// findDigits will check only for digits from 0 to 9
func findDigits(s string) int {
	result := 0
	for _, v := range s {
		if strings.ContainsAny("1234567890", string(v)) {
			result++
		}
	}
	return result
}

// findSymbols will count common symbols appearing on the number keys
func findSymbols(s string) int {
	result := 0
	for _, v := range s {
		if strings.ContainsAny("!@€£#$%^&*()", string(v)) {
			result++
		}
	}
	return result
}

// findWords will count the number of words without surrounding spaces
func findWords(s string) int {
	return len(strings.Fields(s))
}

// printReport will display a summary of results
func printReport(vowels, digits, symbols, words int) {
	fmt.Println("Analysis Summary")
	fmt.Println("-------- -------")
	fmt.Println("vowels  :", vowels)
	fmt.Println("digits  :", digits)
	fmt.Println("symbols :", symbols)
	fmt.Println("words   :", words)
}
