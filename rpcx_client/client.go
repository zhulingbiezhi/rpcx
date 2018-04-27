package main

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/client"
	"rpcx/rpcx_server/service"
)

func main() {
	//op := client.Option{
	//	SerializeType:  protocol.JSON,
	//	Retries:        0,
	//	ConnectTimeout: time.Second * 10,
	//	ReadTimeout:    time.Minute,
	//	WriteTimeout:   time.Second * 10,
	//}

	c := client.NewClient(client.DefaultOption)
	err := c.Connect("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	createReq := &service.CreateRequest{
		Info: "test",
	}
	createResp := &service.CreateResponse{}

	err = c.Call(context.Background(), "PaymentService", "Create", createReq, createResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(createResp)
	select {}
}
