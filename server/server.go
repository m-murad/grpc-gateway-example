package main

import (
	calculatorpb "calculator-service/calculatorpb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	gRPCPort = ":8080"
)

func main() {
	listener, _ := net.Listen("tcp", gRPCPort)
	gRPCServer := grpc.NewServer()

	var server calculatorService
	calculatorpb.RegisterCalculatorServiceServer(gRPCServer, server)

	gRPCServer.Serve(listener)
	defer gRPCServer.GracefulStop()
}

type calculatorService struct{}

func (c calculatorService) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	number1 := req.GetNumber_1()
	number2 := req.GetNumber_2()
	log.Printf("Calculating sum of %d and %d", number1, number2)
	sum := number1 + number2
	response := calculatorpb.SumResponse{
		Sum: sum,
	}
	return &response, nil
}
