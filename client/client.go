package main

import (
	calculatorpb "calculator-service/calculatorpb"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	gRPCPort = ":8080"
)

func main() {
	conn, err := grpc.Dial(gRPCPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()
	client := calculatorpb.NewCalculatorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sumRequest := calculatorpb.SumRequest{
		Number_1: 5,
		Number_2: 7,
	}
	r, err := client.Sum(ctx, &sumRequest)
	if err != nil {
		log.Fatalf("could not calculate sum: %v", err)
	}
	log.Printf("Sum of numbers 5 and 7 is : %d", r.GetSum())
}
