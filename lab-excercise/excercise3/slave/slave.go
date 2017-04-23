package main

import (
	"github.com/jerana/lab-excercise/lab-excercise/excercise3/common"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	w := new(common.WordCnt)
	server := rpc.NewServer()
	server.Register(w)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	//Start listening for message on port 1234
	listener, e := net.Listen("tcp", ":8088")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}
	log.Println("Serving RPC handler")
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error:" + err.Error())
		} else {
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}

}
