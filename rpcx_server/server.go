package main

import (
	"log"

	"github.com/smallnest/rpcx/server"
	"rpcx/rpcx_server/service"
)

func main() {
	serv := server.NewServer()
	serv.RegisterName("PaymentService", new(service.Service), "")
	log.Println("service listen on 8081")
	serv.Serve("tcp", ":8081")
}
