package common

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
)

//Holds arguments to be passed to service wordCount in RPC call
type ReqArgs struct {
	Url []string
}
type Result struct {
	Word  string
	Count int
}
type RespResult []Result

//Representss service wordCount with method  WordCount
type WordCnt struct{}

func (t *WordCnt) WordCount(args ReqArgs, resp *RespResult) error {
	fmt.Println("Recv client req :", len(args.Url))
	tMap := make(map[string]int, 0)
	for _, f := range args.Url {
		fmt.Println("Requested :", f)
		ReadFile(f, tMap)
	}

	*resp = make([]Result, 0)
	for k, v := range tMap {
		*resp = append(*resp, Result{Word: k, Count: v})
	}
	return nil

}
func ReadFile(input string, textMap map[string]int) error {
	resp, err := http.Get(input)
	if err != nil {
		return errors.New("Input file doesn't exist")
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		textMap[scanner.Text()]++
	}
	return nil
}
