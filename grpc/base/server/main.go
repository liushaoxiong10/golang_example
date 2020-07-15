package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/liushaoxiong10/golang_example/grpc/base/pb"

	"google.golang.org/grpc"
)

type UserInfo struct {
}

func (u *UserInfo) GetUserMessage(ctx context.Context, user *pb.UserRequest) (*pb.UserResponse, error) {
	resp := &pb.UserResponse{
		Message: fmt.Sprintf("name: %s, id: %d, phone_number: %d", user.Name, user.Id, user.PhoneNumber),
	}
	return resp, nil
}

func main() {
	l, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatalf("failed to listen： %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserInfoServer(s, new(UserInfo))
	s.Serve(l)
}
