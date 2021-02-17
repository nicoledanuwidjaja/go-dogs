package main

import (
	"context"
	"time"

	"github.com/DataDog/dd-go/apps/reese-counter/grpc"
)

func client() {
	// call gRPC server
	conn, err := pb.Dial("service-name", "127.0.0.1:6401", grpc.WithInsecure())
	client := pb.NewFilesClient(conn)

	stat, err := client.StatFile(ctx, &pb.StatFileRequest{
		Path: remote,
	})

	// contexts create reuseable timeouts/deadlines, cancel ongoing calls (forking, concurrency)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
}
