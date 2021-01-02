package main

import (
	"context"
	"log"

	"github.com/rohit-tambe/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:40019", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection fail %v ", err)
	}
	defer conn.Close()
	c := calculatorpb.NewCalculatorServiceClient(conn)
	unaryCalculatorCall(c)

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
	log.Println("Result ", res.Result)
}
