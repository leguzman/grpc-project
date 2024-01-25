package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/leguzman/grpc-project/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetingServiceClient, namesList *pb.NamesList) {
	log.Printf("Bidirectional Streaming...")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Couldn't setup bidirectional stream: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Error while server streaming: %v", err)
			}
			log.Println(res)
		}
	}()

	for _, name := range namesList.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while client streaming: %v", err)
		}
		log.Printf("Sent %v", name)
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Finished bidirectional streaming")
}
