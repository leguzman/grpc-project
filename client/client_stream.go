package main

import (
	"context"
	"log"
	"time"

	pb "github.com/leguzman/grpc-project/proto"
)

func callSayHelloClientStreaming(client pb.GreetingServiceClient, namesList *pb.NamesList) {
	log.Printf("Client Streaming...")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Couldn't send names: %v", err)
	}

	for _, name := range namesList.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Printf("Sent %v", name)
		time.Sleep(2 * time.Second)

	}
	res, err := stream.CloseAndRecv()
	log.Println("Streaming finished.")
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}
	log.Printf("Messages: %v", res.Messages)
}
