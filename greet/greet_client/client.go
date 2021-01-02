package main

import (
	"context"
	"io"
	"log"

	"github.com/rohit-tambe/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	// unaryGrpcCall(c)
	serverStringGrpcCall(c)

}

func serverStringGrpcCall(c greetpb.GreetServiceClient) {
	resultStream, err := c.GreetManyTime(context.Background(), &greetpb.GreetManyTimeRequest{Name: "Rohit", NumberRequest: 5})
	if err != nil {
		log.Printf("error from greet many time service response %v", err)
	}
	for {
		msg, err := resultStream.Recv()
		if err == io.EOF {
			log.Printf("we get all result %v", err)
			return
		}
		log.Printf("Greet %v ", msg.Result)
	}
}

func unaryGrpcCall(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{
		FirstName: "Rohit",
		LastName:  "Tambe",
	}}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Printf("error from greet response %v", res)
	}
	log.Println(res.Result)
}
