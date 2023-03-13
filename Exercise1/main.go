package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inputFilename := "Exercise1/input.txt"
	outputFilename := "Exercise1/output.txt"

	content, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Fields(string(content))
	numWords := len(words)

	outputContent := fmt.Sprintf("Number of words: %d", numWords)
	err = ioutil.WriteFile(outputFilename, []byte(outputContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of words in %s: %d\n", inputFilename, numWords)
	fmt.Printf("Output written to %s\n", outputFilename)
}
