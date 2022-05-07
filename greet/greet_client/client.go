package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/DJANGO-JANE/GogRPC/greet/greetpb"
	__ "github.com/DJANGO-JANE/GogRPC/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}

	defer con.Close()
	c := __.NewGreetServiceClient(con)
	fmt.Printf("Created client connection %v", c)

	doStreaming(c)
}

// func doUnary(c __.GreetServiceClient) {
// 	req := &__.GreetRequest{
// 		Greeting: &__.Greeting{

// 			FirstName: "Django",
// 			LastName:  "Jane",
// 		},
// 	}
// 	res, err := c.Greet(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error while calling Greet RPC: %v", err)
// 	}
// 	log.Printf("Response : %v", res.Result)
// }

func doStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Starting to do streaming RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Django",
			LastName:  "Jane",
		},
	}
	reqStream, err := client.GreetManyTimes(context.Background(), (*__.GreetManyTimesRequest)(req))
	if err != nil {
		log.Fatalf("Error when calling streaming rpc")
	}

	for {
		msg, err := reqStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream from server.: %v", err)
		}
		log.Printf("Stream : %v", msg.GetResult())
	}
}
