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

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = "8080"
	httpPort = "8081"
)

type server struct {
	pb.UnimplementedUserAuthenticationServer
}

func main() {
	// create server
	server, err := apipkg.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	go runGrpcGatewayServer(server)
	runGrpcServer(server)
}

func runGrpcGatewayServer(server *apipkg.Server) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// mux
	mux := runtime.NewServeMux()

	// register
	pb.RegisterUserAuthenticationHandlerServer(ctx, mux, server)

	// http server
	log.Printf("grpc-gateway server started on localhost:%s", httpPort)
	err := http.ListenAndServe("localhost:"+httpPort, mux)
	if err != nil {
		log.Fatalln(err)
	}
}

func runGrpcServer(server *apipkg.Server) {
	listener, err := net.Listen("tcp", "localhost:"+grpcPort)
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
	log.Printf("grpc server started on localhost:%s", grpcPort)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
