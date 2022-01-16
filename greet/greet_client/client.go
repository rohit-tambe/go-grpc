package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	uneryChan := make(chan bool)
	serverChan := make(chan bool)
	clientChan := make(chan bool)
	go unaryGrpcCall(c, uneryChan)
	go serverStreamGrpcCall(c, serverChan)
	go clientStreaming(c, clientChan)
	<-uneryChan
	<-serverChan
	<-clientChan
}

func serverStreamGrpcCall(c greetpb.GreetServiceClient, ch chan bool) {
	resultStream, err := c.GreetManyTime(context.Background(), &greetpb.GreetManyTimeRequest{Name: "Rohit", NumberRequest: 5})
	if err != nil {
		log.Printf("error from greet many time service response %v", err)
	}
	for {
		msg, err := resultStream.Recv()
		if err == io.EOF {
			log.Printf("we get all result %v", err)
			ch <- true
			return
		}
		log.Printf("Result from Server Streaming Greet %v ", msg.GetResult())
	}

}

func clientStreaming(c greetpb.GreetServiceClient, ch chan bool) {
	// (GreetService_LongGreetClient, error)
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Printf("error from Long greet response %v", err)
	}
	requests := []*greetpb.GreetLongRequest{
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{FirstName: "Rohit", LastName: "Tambe"},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{FirstName: "Ashish", LastName: "Tambe"},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{FirstName: "Holand", LastName: "Tambe"},
		},
	}
	for _, request := range requests {
		stream.Send(request)
		time.Sleep(time.Second * 2)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("error from Long greet CloseAndRecv %v", err)
	}
	fmt.Println("Result from Client Streaming ", res.GetResult())
	ch <- true
}

func unaryGrpcCall(c greetpb.GreetServiceClient, ch chan bool) {
	req := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{
		FirstName: "Rohit",
		LastName:  "Tambe",
	}}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Printf("error from greet response %v", res)
	}
	log.Println("Result from Unery Client ", res.Result)
	ch <- true
}
