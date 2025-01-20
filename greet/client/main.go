package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/wassup-chicken/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change to false if needed

	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//Unary
	// doGreet(c)

	//Stream from Server
	// doGreetManyTimes(c)

	//Stream from Client
	// doLongGreet(c)

	//Bidirectional
	// doGreetEveryone(c)

	//Deadline
	//doGreetWithDeadline(c, 1*time.Second)

	doGreet(c)

}
