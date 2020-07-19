package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/liushaoxiong10/golang_example/grpc/stream/pb"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterLocalTranslateServer(s, &LocalTranslate{})
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

type LocalTranslate struct{}

func (l *LocalTranslate) LocalRotateToName(stream pb.LocalTranslate_LocalRotateToNameServer) error {
	start := time.Now()
	var name []string
	for {
		rota, err := stream.Recv()
		if err == io.EOF {
			use := time.Now().Sub(start)
			name = append(name, fmt.Sprintf("use %s s", use.String()))
			names := strings.Join(name, ",")
			return stream.SendAndClose(&pb.LocalName{Name: names})
		}
		if err != nil {
			return err
		}
		name = append(name, fmt.Sprintf("XX%d:%d", rota.X, rota.Y))
	}
}

func (l *LocalTranslate) GetMore(stream pb.LocalTranslate_GetMoreServer) error {
	start := time.Now()
	for {
		rota, err := stream.Recv()
		if err == io.EOF {
			use := time.Now().Sub(start)
			return stream.Send(&pb.LocalName{Name: fmt.Sprintf("use %s s", use.String())})
		}
		if rota == nil {
			log.Printf("rote nil")
			continue
		}
		if err := stream.Send(&pb.LocalName{Name: fmt.Sprintf("XX%d:%d", rota.X, rota.Y)}); err != nil {
			log.Fatal(err)
		}
	}
}
