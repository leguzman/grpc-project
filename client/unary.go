package main

import (
	"context"
	"log"
	"time"

	pb "github.com/leguzman/grpc-project/proto"
)

func callSayHello(client pb.GreetingServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Couldn't greet: %v ", err)
	}
	log.Printf("%s", res.Message)
}
