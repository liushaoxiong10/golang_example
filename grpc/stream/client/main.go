package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/liushaoxiong10/golang_example/grpc/stream/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8099", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLocalTranslateClient(conn)

	rote := &pb.LocalRotate{}
	log.Printf("begin local stream")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.LocalRotateToName(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		rote.X = int32(i)
		rote.Y = rote.X * 2
		if err := stream.Send(rote); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, rote, err)
		}
	}
	name, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv")
	}
	fmt.Println(name.Name)

	log.Print("local over")

	// double steam
	log.Print("get more begin")
	morestream, err := client.GetMore(ctx)
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := morestream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Get name %s ", in.Name)
		}
	}()
	for i := 0; i < 10; i++ {
		rote.X = int32(i)
		rote.Y = rote.X * 2
		if err := morestream.Send(rote); err != nil {
			log.Fatalf("%v.Send(%v) = %v", morestream, rote, err)
		}
	}
	morestream.CloseSend()
	<-waitc
}
