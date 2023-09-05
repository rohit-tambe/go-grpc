// go:build tool

package mockgen_demo

import (
	_ "github.com/golang/mock/mockgen"
)

//go:generate mockgen -destination=mocks/greetmock/mock_greet.go -package=mocks github.com/rohit-tambe/go-grpc/greetpb/greet_grpc.pb.go GreetServiceServer
