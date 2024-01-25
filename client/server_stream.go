package main

import (
	"context"
	"log"

	pb "github.com/leguzman/grpc-project/proto"
)

func callSayHelloServerStreaming(client pb.GreetingServiceClient, namesList *pb.NamesList) {
	log.Printf("Streaming...")
	stream, err := client.SayHelloServerStreaming(context.Background(), namesList)
	if err != nil {
		log.Fatalf("Couldn't send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err != nil {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Println("Streaming finished.")
}
