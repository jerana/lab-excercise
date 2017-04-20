package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
	"sync"
)

//declare slice of maps to store multiple files data
var data []map[string]int

func main() {
	var wg sync.WaitGroup
	var result map[string]int
	url := "http://www.gutenberg.org/files/15/text/"
	inputFiles, err := readNumfilesFromLink(url)
	if err != nil {
		fmt.Print(err)
		return
	}
	//Init Slice of Maps And let call go routine allocate hashMap for each file
	data = make([]map[string]int, len(inputFiles))

	//Configure Wait group for main thread
	wg.Add(len(inputFiles))

	for idx, f := range inputFiles {
		finalInput := strings.Join([]string{url, f}, "/")
		fmt.Println(idx, ":", finalInput)
		go readInputURLFiles(finalInput, idx, &wg)

	}
	// All go routine must have done by done
	wg.Wait()

	//Allocate hashMap to merge all files hashMap data
	result = make(map[string]int, 0)

	//Walk into slice of hashMap and merge them into final result map
	for _, m := range data {
		//Walk on every map which build during Files parse
		for w, c := range m {
			result[w] += c
		}
	}
	//Write the final result into Output files
	outFile, err := os.Create("excercise2-output.txt")
	if err != nil {
		fmt.Println("failed to create output file")
		return
	}
	defer outFile.Close()
	for w, c := range result {
		fmt.Println(w, c)
		fmt.Fprintf(outFile, "%v %v \n", w, c)
	}

}
func readInputURLFiles(url string, index int, wg *sync.WaitGroup) error {
	resp, err := http.Get(url)
	defer wg.Done()
	if err != nil {
		return errors.New("Invalid URL input file link")

	}
	b := resp.Body
	defer b.Close()
	//Setup new Scanner for given Link file
	scanner := bufio.NewScanner(b)
	//Setup Split func to tokenize input ; Use bufio.ScanWords
	scanner.Split(bufio.ScanWords)
	//Allocat map for given index files,init Map for given file index
	data[index] = make(map[string]int, 0)

	if data[index] == nil {
		err := fmt.Errorf("hash Map is not allocated for index :%d file", index)
		return err
	}
	for scanner.Scan() {
		//Updat hashMap for Given File Index
		w := scanner.Text()
		data[index][w]++
	}
	return nil

}
func readNumfilesFromLink(url string) ([]string, error) {
	var readFiles []string

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Invaild URL link")
	}
	b := resp.Body
	defer b.Close()

	//Use Html package to parse the given URL data
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			fmt.Println("End of Document")
			break
		} else if tt == html.StartTagToken {
			t := z.Token()
			for _, a := range t.Attr {
				if a.Key == "href" {
					readFiles = append(readFiles, a.Val)
				}
			}
		}
	}
	return readFiles, nil
}
