package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/parezban/grpc-go/api/notification"
	"google.golang.org/grpc"
)

var data []Notification

type Notification struct {
	Id      int32
	Message string
	Action  notification.NewDonutNotificationResponse_Action
}

type server struct {
	notification.UnimplementedDonutsNotifierServer
}

func main() {
	fmt.Println("Server is running...")

	// Make a listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Make a gRPC server
	grpcServer := grpc.NewServer()

	notification.RegisterDonutsNotifierServer(grpcServer, &server{})
	// Run the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) NewDonutNotification(ctx context.Context, req *notification.NewDonutNotificationRequest) (*notification.NewDonutNotificationResponse, error) {
	fmt.Printf("Received Sum RPC: %v", req)

	res := &notification.NewDonutNotificationResponse{}

	return res, nil
}
