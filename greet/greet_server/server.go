package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/DJANGO-JANE/GogRPC/greet/greetpb"
	__ "github.com/DJANGO-JANE/GogRPC/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	__.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *__.GreetRequest) (*__.GreetResponse, error) {
	firstname := req.GetGreeting().FirstName
	result := "Hello " + firstname
	res := &__.GreetResponse{
		Result: result,
	}
	return res, nil

}

func (*server) GreetManyTimes(request *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstname := request.GetGreeting().FirstName

	for i := 0; i < 10; i++ {
		result := "Hello " + firstname + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	result := ""
	for {
		request, err := stream.Recv()
		firstname := request.GetGreeting().GetFirstName()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream. %v", err)
		}

		result += firstname + " ! "
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {

	fmt.Println("Incoming bi-directional request")
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("There was an error when reading client stream. %v", err)
		}
		firstname := request.GetGreeting().GetFirstName()
		result := "Hello " + firstname + "! "
		sendError := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendError != nil {
			log.Fatalf("There was an error when sending response to client. %v", err)
			return err
		}
	}
}
func main() {

	fmt.Println("Hello from the server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("There was an error when starting the listener: %v", err)
	}

	defer listener.Close()
	s := grpc.NewServer()
	__.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve. : %v", err)
	}

}
