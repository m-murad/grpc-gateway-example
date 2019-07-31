package main

import (
	calculatorpb "calculator-service/calculatorpb"
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gRPCPort = ":8080"
	httpPort = ":8081"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	calculatorpb.RegisterCalculatorServiceHandlerFromEndpoint(ctx, mux, gRPCPort, opts)
	http.ListenAndServe(httpPort, mux)
}
