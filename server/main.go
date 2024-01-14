// Package main is the entry point for the application.
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	pb "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	apipkg "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/api/userauth"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	grpcPort := getEnvOrDefault("GRPC_PORT", "8080")
	httpPort := getEnvOrDefault("HTTP_PORT", "8081")

	// create server
	server, err := apipkg.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	go runGrpcGatewayServer(server, httpPort)
	runGrpcServer(server, grpcPort)
}

func runGrpcGatewayServer(server *apipkg.Server, httpPort string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// mux
	mux := runtime.NewServeMux()

	// register
	pb.RegisterUserAuthenticationHandlerServer(ctx, mux, server)

	// http server
	log.Printf("grpc-gateway server started on %s:%s", getHost(), httpPort)
	err := http.ListenAndServe(getHost()+":"+httpPort, mux)
	if err != nil {
		log.Fatalln(err)
	}
}

func runGrpcServer(server *apipkg.Server, grpcPort string) {
	listener, err := net.Listen("tcp", getHost()+":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserAuthenticationServer(s, server)
	envType := os.Getenv("ENV_TYPE")

	// Register reflection service on gRPC server.
	if envType != "prod" {
		reflection.Register(s)
	}
	log.Printf("grpc server started on %s:%s", getHost(), grpcPort)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getHost() string {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	return host
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
