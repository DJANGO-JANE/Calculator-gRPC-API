package main

import (
	"context"
	"log"
	"net"

	"github.com/DJANGO-JANE/GogRPC/Calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
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
