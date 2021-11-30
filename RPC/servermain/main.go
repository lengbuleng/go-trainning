package main

import (
	"RPCserver/RPC/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	StringService := new(server.StringService)
	rpc.Register(StringService)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
