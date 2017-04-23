package main

import (
	"fmt"
	"github.com/jerana/lab-excercise/lab-excercise/excercise3/common"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"strconv"
	"strings"
	"sync"
)

var reply common.RespResult
var url string = "http://www.gutenberg.org/files/15/text"

var rMap []map[string]int
var wg sync.WaitGroup

func main() {

	setFiles, err := common.GetInputFilesSet(url)
	if err != nil {
		log.Fatal("Failed to parse given URL:", url)
	}
	rMap = make([]map[string]int, len(setFiles))

	n := 4 //Need to figure out how to get Slave numbers
	wg.Add(len(setFiles))

	for i, file := range setFiles {
		slave := "excercise3_slave_" + strconv.Itoa(i%n+1) + ":8088"

		fmt.Println("Requested service:", slave)

		go sendReqToSlave(file, slave, i, &wg)

	}
	wg.Wait()

	fo, err := os.Create("excercise3-output.txt")
	if err != nil {
		fmt.Println("Failed to open file")
		return
	}
	defer fo.Close()
	for i, r := range rMap {
		//for k, v := range r {
		//fmt.Fprintf(fo, "%v, %v", k, v)
		//}
		fmt.Println("Resp for index:", i, len(r))
	}

}
func sendReqToSlave(f string, dst string, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	client, err := net.Dial("tcp", dst)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	c := jsonrpc.NewClient(client)
	var args common.ReqArgs

	//Syncronous call
	fileUrl := strings.Join([]string{url, f}, "/")
	args.Url = append(args.Url, fileUrl)
	fmt.Println("Going to write to server")
	err = c.Call("WordCnt.WordCount", args, &reply)
	if err != nil {
		fmt.Println("Failed to connet server")
		log.Fatal("wordcnt fail:", err)
	}
	fmt.Println("Got result for index: ", index)
	fmt.Println(reply)

	rMap[index] = make(map[string]int, 0)

	for _, r := range reply {
		rMap[index][r.Word] += r.Count
	}

}
