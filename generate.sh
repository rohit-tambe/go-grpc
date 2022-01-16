#!/bin/bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
protoc pcbook/proto/processor_message.proto --go_out=plugins=grpc:.