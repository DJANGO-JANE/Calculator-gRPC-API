package main

import (
	"context"
	"fmt"
	"log"

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

	doUnary(c)
}

func doUnary(c __.GreetServiceClient) {
	req := &__.GreetRequest{
		Greeting: &__.Greeting{

			FirstName: "Django",
			LastName:  "Jane",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response : %v", res.Result)
}
