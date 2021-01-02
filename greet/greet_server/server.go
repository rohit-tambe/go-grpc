package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/rohit-tambe/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// Greet(context.Context, *GreetRequest) (*GreetResponse, error)
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("Function invoke ", req.GetGreeting().FirstName)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}
func (*server) GreetManyTime(req *greetpb.GreetManyTimeRequest, stream greetpb.GreetService_GreetManyTimeServer) error {
	for i := 0; i < int(req.NumberRequest); i++ {
		res := &greetpb.GreetManyTimeResponse{
			Result: fmt.Sprintf("Hello %s with number %d", req.Name, i),
		}
		stream.Send(res)
		time.Sleep(time.Second * 5)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	fmt.Println("server starts on ", lis.Addr())
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
