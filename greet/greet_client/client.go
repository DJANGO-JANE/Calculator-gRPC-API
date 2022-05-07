package main

import (
	"fmt"
	"log"

	"github.com/DJANGO-JANE/GogRPC/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")

	con, err := grpc.Dial("localhost:6936", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}

	defer con.Close()
	c := greetpb.NewGreetServiceClient(con)
	fmt.Printf("Created client connection %v", c)
}
