package main

import (
	context "context"

	"google.golang.org/grpc"

	"fmt"
	"log"
	"net"

	myPb "fredericschmidt.fr/grpcdemos/GoDemo01/calculatrice"
)

// Port for the gRPC Server
const (
	port = ":666"
)

type CalculatorServer struct {
	myPb.UnimplementedCalculatorServiceServer
}

// Implementation
func (s *CalculatorServer) AddOperation(ctx context.Context, in *myPb.ValuesCalculatorRequest) (*myPb.ResultCalculatorResponse, error) {
	return &myPb.ResultCalculatorResponse{Result: in.TermX + in.TermY}, nil
}

func (s *CalculatorServer) SubOperation(ctx context.Context, in *myPb.ValuesCalculatorRequest) (*myPb.ResultCalculatorResponse, error) {
	return &myPb.ResultCalculatorResponse{Result: in.TermX - in.TermY}, nil
}

func (s *CalculatorServer) TableOperation(ctx context.Context, in *myPb.TableCalculatorRequest) (*myPb.TableCalculatorResponse, error) {

	r := make([]*myPb.OneLineInTableResponse, 0)

	fmt.Printf("TableOperation Called\n")

	var i int32
	for i = 0; i <= in.Multiplier; i++ {

		r = append(r, &myPb.OneLineInTableResponse{
			Multiplicand: in.Multiplicand,
			Multiplier:   i,
			Product:      i * in.Multiplicand,
		})
	}
	return &myPb.TableCalculatorResponse{LineOfTable: r}, nil
}

// Main Func
func main() {

	fmt.Println("Starting Calculator Server gRPC")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("ERROR : failed to listen on port %s : %v", port, err)
	}

	srv := grpc.NewServer()
	myPb.RegisterCalculatorServiceServer(srv, &CalculatorServer{})
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
