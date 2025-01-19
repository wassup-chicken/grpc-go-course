package main

import (
	"context"
	"log"

	pb "github.com/wassup-chicken/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})

	if err != nil {
		log.Fatalf("Could not grteet: %v\n", err)
	}

	log.Printf("Greeting: %ss\n", res.Result)
}
