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
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	dbStore, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	grpcPort := getEnvOrDefault("GRPC_PORT", "8080")
	httpPort := getEnvOrDefault("PORT", "8081")

	// create server
	server, err := apipkg.NewServer(dbStore)
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
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
			if customErr, ok := status.FromError(err); ok {
				http.Error(w, customErr.Message(), runtime.HTTPStatusFromCode(customErr.Code()))
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}),
	)

	// register
	pb.RegisterUserAuthenticationHandlerServer(ctx, mux, server)

	// http server
	log.Printf("grpc-gateway server started on %s:%s", getEnvOrDefault("HOST", "localhost"), httpPort)
	err := http.ListenAndServe(getEnvOrDefault("HOST", "localhost")+":"+httpPort, mux)
	if err != nil {
		log.Fatalln(err)
	}
}

func runGrpcServer(server *apipkg.Server, grpcPort string) {
	listener, err := net.Listen("tcp", getEnvOrDefault("HOST", "localhost")+":"+grpcPort)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			resp, err := handler(ctx, req)
			if err != nil {
				if customErr, ok := status.FromError(err); ok {
					return nil, customErr.Err()
				}
				return nil, status.Error(codes.Internal, err.Error())
			}
			return resp, nil
		}),
	)

	pb.RegisterUserAuthenticationServer(grpcServer, server)
	reflection.Register(grpcServer)

	log.Printf("grpc server started on %s:%s", getEnvOrDefault("HOST", "localhost"), grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
