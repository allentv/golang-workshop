// Reading data from a file and then print out the summary
// Path to "sample.txt" => src/app/variation/sample.txt relative to project root

package main

import (
	"app/utils"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

// cleanupInput runs all the validation and bleaching to be applied on the input string
func cleanupInput(inputFileName string) string {
	if inputFileName == "" {
		log.Fatal("Input file name not provided!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		log.Fatal("Error => ", err.Error())
		os.Exit(1)
	}

	return string(data)
}

func main() {
	inputFileName := flag.String("file", "", "File to parse")
	flag.Parse()

	cleanedStr := cleanupInput(*inputFileName)

	vowels := utils.FindVowels(cleanedStr)
	digits := utils.FindDigits(cleanedStr)
	symbols := utils.FindSymbols(cleanedStr)
	words := utils.FindWords(cleanedStr)

	utils.PrintReport(vowels, digits, symbols, words)
}
