package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"imooc/rpc"
	"log"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	rpc.Register(rpcdemo.DemoService1{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error :%v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
