package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/rohit-tambe/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

// Greet(context.Context, *GreetRequest) (*GreetResponse, error)
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("Function server streaming invoke ", req.GetGreeting().FirstName)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			//client cancele the request
			log.Println("client cancele the request")
			return nil, status.Error(codes.Canceled, "client cancele the request")
		}
		time.Sleep(time.Second * 1)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

// SayHelloName(context.Context, *GreetRequest) (*GreetResponse, error)
func (*server) SayHelloName(ctx context.Context, req *greetpb.SayHello) (*greetpb.SayHello, error) {
	log.Println("Function server streaming invoke ", req.GetFirstName())
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			//client cancele the request
			log.Println("client cancele the request")
			return nil, status.Error(codes.Canceled, "client cancele the request")
		}
		time.Sleep(time.Second * 1)
	}
	firstName := req.GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.SayHello{
		FirstName: result,
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

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &greetpb.GreetLongResponse{Result: result}
			return stream.SendAndClose(res)
		}
		if err != nil {
			return err
		}
		result += fmt.Sprintf("Hello %s\n", req.GetGreeting().GetFirstName())
	}
}
func (*server) GreetEveryOne(stream greetpb.GreetService_GreetEveryOneServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}
		log.Println("bidi server request ", req.GetGreeting().GetFirstName())
		result := "Hello " + req.GetGreeting().GetFirstName() + " " + req.GetGreeting().GetLastName()
		err = stream.Send(&greetpb.GreetEveryOneResponse{Result: result})
		if err != nil {
			return err
		}
	}
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
