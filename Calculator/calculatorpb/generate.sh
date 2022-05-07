#!/bin/bash

protoc calculator/calculatorpb/calculator.proto --go-grpc_out=.
protoc greet/greetpb/greet.proto --go_out=. --go-grpc_out=.