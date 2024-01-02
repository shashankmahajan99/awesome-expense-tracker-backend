// Package main is the entry point for the application.
package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserAuthenticationServer
}

func main() {
	// create server
	server, err := pkg.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	go runGrpcGatewayServer(server)
	runGrpcServer(server)
}

func runGrpcGatewayServer(server *pkg.Server) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// mux
	mux := runtime.NewServeMux()

	// register
	pb.RegisterUserAuthenticationHandlerServer(ctx, mux, server)

	// http server
	log.Fatalln(http.ListenAndServe(":8081", mux))
}

func runGrpcServer(server *pkg.Server) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserAuthenticationServer(s, server)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
