package main

import (
	"context"
	"fmt"
	"io"
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

	doStreamRequest(client)

}

// func doUnaryRequest(client calculatorpb.CalculatorServiceClient) {
// 	req := &calculatorpb.CalculateRequest{
// 		Calculate: &calculatorpb.Calculate{
// 			FirstNum:  10,
// 			SecondNum: 3,
// 		},
// 	}
// 	res, err := client.Calculate(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("An error occurred when calling Calculate RPC: %v", err)
// 	}
// 	log.Printf("Response : %v", res.Result)
// }

func doStreamRequest(client calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculateManyTimesRequest{
		Calculate: &calculatorpb.Calculate{
			FirstNum: 120,
		},
	}
	reqStream, err := client.CalculateStream(context.Background(), req)
	if err != nil {
		log.Fatalf("An error occurred when streaming Calulate RPC: %v", err)
	}

	for {
		msg, err := reqStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error when reading stream from server. %v", err)
		}
		log.Printf("Stream : %v", msg.GetResult())
	}
}
