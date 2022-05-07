package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	doBiDiStreaming(client)

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

// func doStreamRequest(client calculatorpb.CalculatorServiceClient) {

// 	requests := []*calculatorpb.AverageRequest{
// 		{
// 			Number: 1,
// 		},
// 		{
// 			Number: 2,
// 		},
// 		{
// 			Number: 3,
// 		},
// 		{
// 			Number: 4,
// 		},
// 	}
// 	stream, err := client.ComputeAverage(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error when making client streaming request")
// 	}
// 	for _, req := range requests {
// 		fmt.Printf("Sending requests to server...%v\n", req)
// 		stream.Send(req)
// 		time.Sleep(100 * time.Millisecond)
// 	}

// 	response, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatalf("Error occurred when closing and receiving")
// 	}
// 	fmt.Printf("Response is :%v", response.Result)

// }

// func doClientStreaming(client calculatorpb.CalculatorServiceClient) {
// 	req := &calculatorpb.CalculateManyTimesRequest{
// 		Calculate: &calculatorpb.Calculate{
// 			FirstNum: 120,
// 		},
// 	}
// 	reqStream, err := client.CalculateStream(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("An error occurred when streaming Calulate RPC: %v", err)
// 	}

// 	for {
// 		msg, err := reqStream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Error when reading stream from server. %v", err)
// 		}
// 		log.Printf("Stream : %v", msg.GetResult())
// 	}
// }

func doBiDiStreaming(client calculatorpb.CalculatorServiceClient) {
	log.Println("Making a streaming request to server.")
	requests := []*calculatorpb.MaxNumberRequest{
		{
			Number: 1,
		},
		{
			Number: 5,
		},

		{
			Number: 3,
		},

		{
			Number: 6,
		},
		{
			Number: 2,
		},
		{
			Number: 20,
		},
	}

	stream, err := client.FindMaximum(context.Background())
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
			fmt.Printf("Received response : %v\n", response.Number)
		}
		close(waitchannel)
	}()
	<-waitchannel
}
