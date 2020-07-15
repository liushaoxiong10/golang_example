package main

import (
	"context"
	"fmt"
	"log"

	"github.com/liushaoxiong10/golang_example/grpc/base/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8099", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	userClient := pb.NewUserInfoClient(conn)
	userReq := &pb.UserRequest{
		Name:        "lsx",
		Id:          1,
		PhoneNumber: 12345,
	}
	reply, err := userClient.GetUserMessage(context.Background(), userReq)
	fmt.Println(reply.GetMessage(), err)
}
