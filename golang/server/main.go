package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "proto/proto"
	db "proto/server/config"
	bookstore "proto/server/handlers"
)

var (
	port = ":8000"
)

func startGRPCServer() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}

	s := grpc.NewServer()
	pb.RegisterBookstoreServer(s, &bookstore.Server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	s.Serve(ln)
}

func main() {
	go db.Init()
	startGRPCServer()
}
