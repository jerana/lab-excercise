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

	//Need to set WRK_NODE env in compose run cmd
	n, err := strconv.Atoi(os.Getenv("WRK_NODE"))
	if err != nil {
		fmt.Println("Failed to convert into int:", m)
		return
	}

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
		for k, v := range r {
			fmt.Fprintf(fo, "%v:%v \n", k, v)
		}
		fmt.Println("writing index map:", i)
	}

}
func sendReqToSlave(f string, dst string, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	client, err := net.Dial("tcp", dst)
	if err != nil {
		//fmt.Println("Failing at Dialing:", dst)
		log.Fatal("dialing:", err)
	}
	c := jsonrpc.NewClient(client)
	var args common.ReqArgs

	//Syncronous call
	fileUrl := strings.Join([]string{url, f}, "/")
	args.Url = append(args.Url, fileUrl)
	err = c.Call("WordCnt.WordCount", args, &reply)
	if err != nil {
		log.Fatal("wordcnt fail:", err)
	}

	rMap[index] = make(map[string]int, 0)

	for _, r := range reply {
		rMap[index][r.Word] += r.Count
	}
}
