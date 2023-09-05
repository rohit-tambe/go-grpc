#!/bin/bash
#protoc greet/greetpb/greet.proto --go_out=.
#protoc --proto_path=greet greetpb/greet.proto --go_out=.
protoc --proto_path=greet greetpb/greet.proto --go-grpc_out=.
#protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
#protoc pcbook/proto/processor_message.proto --go_out=plugins=grpc:.
