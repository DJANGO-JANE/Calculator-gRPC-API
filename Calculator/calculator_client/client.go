package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DJANGO-JANE/GogRPC/Calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello from client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to 'localhost:50051': %v", err)
	}

	defer conn.Close()

	client := calculatorpb.NewCalculatorServiceClient(conn)
	fmt.Println("Client connection started.")

	doUnary(client)

}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculateRequest{
		Calculate: &calculatorpb.Calculate{
			FirstNum:  10,
			SecondNum: 3,
		},
	}
	res, err := client.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("An error occurred when calling Calculate RPC: %v", err)
	}
	log.Printf("Response : %v", res.Result)
}
