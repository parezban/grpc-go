package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/parezban/grpc-go/api/notification"
	"google.golang.org/grpc"
)

var data []Notification

type Notification struct {
	Message string
}

type server struct {
	notification.UnimplementedDonutsNotifierServer
}

func main() {
	fmt.Println("Server is running...")

	// Make a listener
	lis, err := net.Listen("tcp", "0.0.0.0:3000")
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

func (*server) GetNotification(req *notification.ListDonutsNotificationsRequest, stream notification.DonutsNotifier_ListDonutsNotificationsServer) error {

	var lastDataPushed []Notification

	for {
		if len(lastDataPushed) != len(data) {
			for _, newData := range data[len(lastDataPushed):] {
				err := stream.Send(&notification.NewDonutNotificationResponse{
					Message: newData.Message,
				})
				if err != nil {
					log.Fatalf("Failed to send response: %v\n", err)
				} else {
					lastDataPushed = append(lastDataPushed, Notification{
						Message: newData.Message,
					})
				}
			}
		}
		time.Sleep(time.Second / 2)
	}
}
