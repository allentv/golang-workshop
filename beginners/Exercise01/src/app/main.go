// Reading data from the command line and then print out the summary

package main

import (
	"app/utils"
	"flag"
	"log"
	"os"
	"strings"
)

// cleanupInput runs all the validation and bleaching to be applied on the input string
func cleanupInput(inputStr string) string {
	if inputStr == "" {
		log.Fatal("Input string not provided!")
		flag.PrintDefaults()
		os.Exit(1)
	}
	return strings.ToLower(inputStr)
}

func main() {
	inputStr := flag.String("text", "", "Text block to parse")
	flag.Parse()

	cleanedStr := cleanupInput(*inputStr)

	vowels := utils.FindVowels(cleanedStr)
	digits := utils.FindDigits(cleanedStr)
	symbols := utils.FindSymbols(cleanedStr)
	words := utils.FindWords(cleanedStr)

	utils.PrintReport(vowels, digits, symbols, words)
}
