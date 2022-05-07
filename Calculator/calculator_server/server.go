package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/DJANGO-JANE/GogRPC/Calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	result := 0.0
	total := 0
	counter := 0
	for {
		request, err := stream.Recv()

		nums := request.GetNumber()
		log.Printf("Number is %v. Counter is :%v", nums, counter)
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error when reading from client stream. : %v", err)
		}
		counter++
		total += int(nums)
		result = float64(total) / float64(counter)
	}
}
func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculateRequest) (*calculatorpb.CalculateResponse, error) {
	num1 := req.GetCalculate().GetFirstNum()
	num2 := req.GetCalculate().GetSecondNum()

	sum := num1 + num2
	res := &calculatorpb.CalculateResponse{
		Result: sum,
	}
	return res, nil
}

func (*server) CalculateStream(request *calculatorpb.CalculateManyTimesRequest, stream calculatorpb.CalculatorService_CalculateStreamServer) error {
	num1 := request.GetCalculate().FirstNum
	var divisor int32
	divisor = 2
	for num1 > 1 {
		if num1%divisor == 0 {
			num1 = num1 / divisor
		} else {
			divisor = divisor + 1
		}
		response := &calculatorpb.CalculateManyTimesResponse{
			Result: divisor,
		}
		stream.Send(response)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil

}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {

	fmt.Println("Incoming Bi-Di request")
	max := int32(0)

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("There was an error when reading from Bi-Di Client.: %v", err)
		}

		//var numArr []int32
		num := request.GetNumber()
		if num > max {
			log.Printf("Maximum is :%d, number is %d", max, num)
			max = num
			sendError := stream.Send(&calculatorpb.MaxNumberResponse{
				Number: int32(max),
			})

			if sendError != nil {
				log.Fatalf("There was an error when responding to Bi-di client")
				return err
			}
		}

	}
}
func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("There was an error when trying to listen")
	} else {

		log.Println("Hello. Listener is running.")
	}

	defer listener.Close()

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Server failed to serve: %v", err)
	}

}
