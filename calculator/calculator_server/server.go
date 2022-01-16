package main

import (
	"context"
	"io"
	"log"
	"net"

	calculatorpb "github.com/rohit-tambe/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumRepsonse, error) {

	out := &calculatorpb.SumRepsonse{
		Result: req.GetFirstNumber() + req.GetSecondNumber(),
	}
	return out, nil
}

// *calculatorpb.PrimeNumberDecompositionRequest
func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	n := req.GetNumber()
	divisor := int32(2)
	for n > 1 {
		if n%divisor == 0 {
			res := &calculatorpb.PrimeNumberDecompositionResponse{Result: divisor}
			stream.Send(res)
			n = n / divisor
		} else {
			divisor = divisor + 1
		}
	}
	return nil
}
func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	sum := int32(0)
	for count := 0; ; count++ {
		request, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			res := &calculatorpb.ComputeAverageResponse{Result: average}
			return stream.SendAndClose(res)
		}
		sum = sum + request.GetNumber()
	}
	return nil
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
