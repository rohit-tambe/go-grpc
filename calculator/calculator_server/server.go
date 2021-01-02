package main

import (
	"context"
	"log"
	"net"

	"github.com/rohit-tambe/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumRepsonse, error) {

	out := &calculatorpb.SumRepsonse{
		Result: req.GetFirstNumber() + req.GetSecondNumber(),
	}
	return out, nil
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:40019")
	if err != nil {
		log.Printf("server dosent listen %v", err)
	}
	log.Printf("server listen %v", lis.Addr())

	sv := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(sv, &server{})
	if err := sv.Serve(lis); err != nil {
		log.Printf("fail to serve %v", err)
	}
}
