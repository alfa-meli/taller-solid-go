package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inputFilename := "./input.txt"
	outputFilename := "./output.txt"

	numWords := count(readFile(inputFilename))

	outputContent := fmt.Sprintf("Number of words: %d", numWords)
	writeFile(outputFilename, outputContent)

	print(inputFilename, outputFilename, numWords)
}

func readFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func count(content string) int {
	words := strings.Fields(content)

	return len(words)
}

func writeFile(filePath, content string) {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func print(in, out string, n int) {
	fmt.Printf("Number of words in %s: %d\n", in, n)
	fmt.Printf("Output written to %s\n", out)
}
