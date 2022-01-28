package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/rohit-tambe/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

var wg sync.WaitGroup

func main() {
	conn, err := grpc.Dial("localhost:40019", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection fail %v ", err)
	}
	defer conn.Close()
	c := calculatorpb.NewCalculatorServiceClient(conn)
	// unaryCalculatorCall(c)
	// primeNumberDecompositionCall(c, 100)
	// computeAverage(c, 1, 2, 3, 4, 5, 6)
	findMaximum(c, 1, 2, 3, 4, 5, 6)
}

func findMaximum(c calculatorpb.CalculatorServiceClient, numbers ...int32) {
	stream, err := c.FindMaximumValue(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for _, value := range numbers {
			err = stream.Send(&calculatorpb.FindMaximumValueRequest{Number: value})
			if err != nil {
				wg.Done()
				log.Println("Send findMaximum err ", err)
				return
			}
			time.Sleep(time.Second * 2)
		}
		err = stream.CloseSend()
		if err != nil {
			wg.Done()
			log.Println("CloseSend findMaximum err ", err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Recv findMaximum  ", res.Result)
				wg.Done()
				return
			}
			if err != nil {
				wg.Done()
				log.Println("Recv findMaximum err ", err)
				return
			}
		}
	}(&wg)
	wg.Wait()
}

func computeAverage(c calculatorpb.CalculatorServiceClient, numbers ...int32) {
	// ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error from ComputeAverage service %v ", err)
	}
	for _, number := range numbers {
		req := &calculatorpb.ComputeAverageRequest{Number: number}
		stream.Send(req)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error from ComputeAverage service %v ", err)
	}
	fmt.Println("Result from Client stream ComputeAverage ", response.GetResult())
}

func primeNumberDecompositionCall(c calculatorpb.CalculatorServiceClient, i int) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{Number: int32(2345654)}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error from PrimeNumberDecomposition service %v ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Printf("we get all result %v", err)
			return
		}
		log.Printf("Server streaming Prime Factor %v ", res.GetResult())
	}
}

func unaryCalculatorCall(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber:  15,
		SecondNumber: 6,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error from Sum service %v ", err)
	}
	log.Println("Sum unery call Result ", res.Result)
}
