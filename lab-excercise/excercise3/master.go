package main

import (
	"fmt"
	"github.com/jerana/lab-excercise/excercise3/common"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

var reply common.RespResult

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	rMap := make(map[string]int, 0)
	//Syncronous call
	var args common.ReqArgs
	args.Url = append(args.Url, "http://www.gutenberg.org/files/15/text/moby-000.txt")
	c := jsonrpc.NewClient(client)
	fmt.Println("Going to call RPC server")
	err = c.Call("WordCnt.WordCount", args, &reply)
	if err != nil {
		fmt.Println("Failed to connet server")
		log.Fatal("wordcnt fail:", err)
	}
	fmt.Println("Got result ", reply)
	for _, r := range reply {
		rMap[r.Word] += r.Count
	}
	for k, v := range rMap {
		fmt.Println(k, v)
	}

}
