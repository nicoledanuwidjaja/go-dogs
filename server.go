package main

import (
	"net"

	"github.com/DataDog/dd-go/apps/reese-counter/grpc"
)

func server() {
	li, err := net.Listen("tcp", "127.0.0.1:6400")
	if err != nil {
		fatalf("failed to start server: %v", err)
	}
	defer li.Close()

	s := grpc.NewServer()
	// attach file services to server
	pb.RegisterFilesServer(s, NewServer())
}
