package main

import (
	"context"
	"fmt"

	"github.com/liushaoxiong10/golang_example/rpcx/base/model"
	"github.com/smallnest/rpcx/client"
)

func main() {

	// #1
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8099", "")
	// #2
	xclient := client.NewXClient("User", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3
	args := &model.UserInfo{
		Name: "lsx",
		ID:   10,
	}

	// #4
	reply := &model.UserReply{}

	// #5
	err := xclient.Call(context.Background(), "GetInfo", args, reply)
	if err != nil {
		fmt.Printf("failed to call: %v", err)
	}
	fmt.Println(reply.Message)
}
