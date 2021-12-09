package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const portNumber = "8000" // port 번호

func main() {
	listen, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("Start GRPC server on %s port", portNumber)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to Serve : %s", err)
	}
}
