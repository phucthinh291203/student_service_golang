package main

import (
	database "grpc-project/database"
	pb "grpc-project/proto"
	services "grpc-project/services"
	"log"
	"net"

	"google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

func main() {
	database.ConnectToDatabase()
	grpcServer := grpc.NewServer()
	studentService := services.NewStudentService()
	pb.RegisterStudentServiceServer(grpcServer, studentService)

	reflection.Register(grpcServer)
	// Lắng nghe trên cổng 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
