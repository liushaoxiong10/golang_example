package main

import (
	"github.com/liushaoxiong10/golang_example/rpcx/base/model"
	"github.com/smallnest/rpcx/server"
)

func main() {
	s := server.NewServer()
	// s.RegisterName("User", new(model.User), "")
	s.Register(new(model.User), "")
	s.Serve("tcp", ":8099")
}
