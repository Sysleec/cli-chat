package root

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectChatServer() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect chat server: %v", err)
	}
	return conn
}

func ConnectLoginServer() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50061", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect login server: %v", err)
	}
	return conn
}
