package main

import (
	"net/rpc"
	"net"
	"log"
	"net/http"
	"server"
)

func main() {
	arith := new(rpcs.Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil)
}
