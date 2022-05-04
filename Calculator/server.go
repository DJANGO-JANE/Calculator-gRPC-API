package main

import (
	"context"
	"net"

	"github.com/DJANGO-JANE/GogRPC/Calculator/calculatorpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server CalculatorServiceServer{}

func (*server) Calculate(ctx context.Context, request *calculatorpb.CalculateRequest) (*CalculateResponse, error) {
	request
}

func main() {

	listening, err := net.Listen("tcp", "0.0.0.0:5050")
	if err != nil{
		log.WithFields(log.Fields{
			"error":err,
		}).Error("There was an error when trying to listen")
	}else{

		log.Info("Hello. Listener is running.")
	}
	
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s,)

}
