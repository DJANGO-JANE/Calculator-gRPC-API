package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	doBiDiStreaming(c)
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

// func doStreaming(client greetpb.GreetServiceClient) {
// 	fmt.Println("Starting to do streaming RPC...")

// 	req := &greetpb.GreetRequest{
// 		Greeting: &greetpb.Greeting{
// 			FirstName: "Django",
// 			LastName:  "Jane",
// 		},
// 	}
// 	reqStream, err := client.GreetManyTimes(context.Background(), (*__.GreetManyTimesRequest)(req))
// 	if err != nil {
// 		log.Fatalf("Error when calling streaming rpc")
// 	}

// 	for {
// 		msg, err := reqStream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Error while reading stream from server.: %v", err)
// 		}
// 		log.Printf("Stream : %v", msg.GetResult())
// 	}
// }

// func doClientStreaming(client greetpb.GreetServiceClient) {

// 	requests := []*greetpb.LongGreetRequest{
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Panther",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Lion",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Bear",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Fox",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Wolf",
// 			},
// 		},
// 	}

// 	stream, err := client.LongGreet(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error when making client stream request. %v", err)
// 	}

// 	for _, req := range requests {
// 		fmt.Printf("Sending requests: %v\n", req)
// 		stream.Send(req)
// 		time.Sleep(100 * time.Millisecond)
// 	}

// 	response, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatalf("Error occured when closing and receiving")
// 	}

// 	fmt.Printf("Response from server : %v", response)
// }

func doBiDiStreaming(client greetpb.GreetServiceClient) {
	log.Println("Making a streaming request to server.")
	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Panther",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lion",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Bear",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Fox",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Wolf",
			},
		},
	}

	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error when sending stream.")
	}

	waitchannel := make(chan struct{})
	//Stream a stream of messages
	go func() {
		for _, item := range requests {
			fmt.Printf("Sending greeting stream: %v\n", item)
			stream.Send(item)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {

			response, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving stream from server. %v", err)
				break
			}
			fmt.Printf("Received response : %v\n", response.Result)
		}
	}()
	<-waitchannel
}
