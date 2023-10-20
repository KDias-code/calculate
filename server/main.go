package main

import (
	"fmt"
	__ "github.com/KDias-code/calculate/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	__.RegisterCalculatorServer(srv, &server{})

	fmt.Println("Server is running on :50051...")
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
