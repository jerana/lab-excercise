package main

import (
	"fmt"
	"github.com/jerana/lab-excercise/lab-excercise/excercise3/common"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"strings"
)

var reply common.RespResult
var url string = "http://www.gutenberg.org/files/15/text"

func main() {
	var args common.ReqArgs
	rMap := make(map[string]int, 0)

	setFiles, err := common.GetInputFilesSet(url)
	if err != nil {
		log.Fatal("Failed to parse given URL:", url)
	}

	client, err := net.Dial("tcp", "slave")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	c := jsonrpc.NewClient(client)

	//Syncronous call
	for _, f := range setFiles {
		fileUrl := strings.Join([]string{url, f}, "/")
		args.Url = append(args.Url, fileUrl)
		err = c.Call("WordCnt.WordCount", args, &reply)
		if err != nil {
			fmt.Println("Failed to connet server")
			log.Fatal("wordcnt fail:", err)
		}
		fmt.Println("Got result ", reply)
		for _, r := range reply {
			rMap[r.Word] += r.Count
		}
	}
	//Write into file
	fo, err := os.Create("excercise-output.txt")
	if err != nil {
		fmt.Println("Failed to open file")
	}
	defer fo.Close()
	for k, v := range rMap {
		fmt.Fprintf(fo, "%v, %v", k, v)
	}

}
