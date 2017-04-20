package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
)

var textMap map[string]int

func main() {
	fmt.Println("vim-go")
	url := "http://www.gutenberg.org/files/15/text/moby-000.txt"
	err := seqReadFile(url)
	if err != nil {
		return
	}
	outFile, err := os.Create("project1-output.txt")
	if err != nil {
		fmt.Println("Failed to create output file!!!")
		return
	}
	defer outFile.Close()
	for words, count := range textMap {
		fmt.Println(words, " ", count)
		fmt.Fprintf(outFile, "%v %v\n", words, count)
	}
}

func seqReadFile(input string) error {
	resp, err := http.Get(input)
	if err != nil {
		return errors.New("Input file doesn't exist")
	}
	defer resp.Body.Close()

	//Setup New Scanner for given file link
	scanner := bufio.NewScanner(resp.Body)
	//Setup Split function to tokenize the input ; use bufio.ScanWords
	scanner.Split(bufio.ScanWords)

	textMap = make(map[string]int, 0)
	for scanner.Scan() {
		textMap[scanner.Text()]++
	}
	return nil
}
