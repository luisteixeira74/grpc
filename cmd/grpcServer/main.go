package main

import (
	"database/sql"
	"net"

	"github.com/luisteixeira74/grpc/internal/database"
	"github.com/luisteixeira74/grpc/internal/pb"
	"github.com/luisteixeira74/grpc/internal/service"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)
	grpcServer := grpc.NewServer()

	// Register the service with the gRPC server
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// Register reflection service on gRPC server
	// This is useful for debugging and testing
	// It allows clients to discover the services and methods available on the server
	// without having to know the service name or method name in advance
	// It is not recommended to use reflection in production code
	reflection.Register(grpcServer)

	// Open a TCP connection on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
	// Close the listener when the server stops
	defer listener.Close()
}
