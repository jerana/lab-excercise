package common

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func GetInputFilesSet(url string) ([]string, error) {
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
