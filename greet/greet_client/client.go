package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/rohit-tambe/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var wg sync.WaitGroup

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	uneryChan := make(chan bool)
	// serverChan := make(chan bool)
	// clientChan := make(chan bool)
	// bidiChan := make(chan bool)
	go unaryGrpcCall(c, uneryChan, 5)
	go unaryGrpcCall(c, uneryChan, 1)
	// go serverStreamGrpcCall(c, serverChan)
	// go clientStreaming(c, clientChan)
	// go bidiStreaming(c, bidiChan)
	<-uneryChan
	// <-serverChan
	// <-clientChan
	// <-bidiChan
}

func bidiStreaming(c greetpb.GreetServiceClient, bidiChan chan bool) {
	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		bidiChan <- true
		log.Printf("error from greet many time service response %v", err)
	}
	requests := []*greetpb.GreetEveryOneRequest{
		{
			Greeting: &greetpb.Greeting{FirstName: "Sam", LastName: "BiDi"},
		},
		{
			Greeting: &greetpb.Greeting{FirstName: "Alexa", LastName: "BiDi"},
		},
		{
			Greeting: &greetpb.Greeting{FirstName: "Holand", LastName: "BiDi"},
		},
		{
			Greeting: &greetpb.Greeting{FirstName: "Tom", LastName: "BiDi"},
		},
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {

		for _, req := range requests {
			err = stream.Send(req)
			time.Sleep(time.Second * 2)
			if err != nil {
				log.Printf("send error from greet bidi many time service response %v", err)
				break
			}
		}
		err = stream.CloseSend()
		if err != nil {
			log.Printf("CloseSend error from greet many bidi time service response %v", err)
		}
	}(&wg)
	func(wg *sync.WaitGroup) {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				break
			}
			if err != nil {
				wg.Done()
				log.Printf("Recve rror from greet many time bidi service response %v", err)
				break
			}
			fmt.Println("Recieved BIDI streaming result ", res.GetResult())
		}
	}(&wg)
	bidiChan <- true
	wg.Wait()
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
		{
			Greeting: &greetpb.Greeting{FirstName: "Rohit", LastName: "Tambe"},
		},
		{
			Greeting: &greetpb.Greeting{FirstName: "Ashish", LastName: "Tambe"},
		},
		{
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

func unaryGrpcCall(c greetpb.GreetServiceClient, ch chan bool, timeout time.Duration) {
	req := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{
		FirstName: "Rohit",
		LastName:  "Tambe",
	}}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*timeout)
	defer cancel()
	res, err := c.Greet(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("time out is hit !deadline exeeded", statusErr.Code(), " with time ", timeout)
			}
			return
		}
		log.Printf("error from greet response %v", statusErr)
		return
	}
	log.Println("Result from Unery Client ", res.Result)
	ch <- true
}
