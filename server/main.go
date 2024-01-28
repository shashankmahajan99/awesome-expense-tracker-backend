// Package main is the entry point for the application.
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	pb "github.com/shashankmahajan99/awesome-expense-tracker-backend/api"
	apipkg "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/api"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/api/auth"
	"github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/api/middlewares"
	db "github.com/shashankmahajan99/awesome-expense-tracker-backend/pkg/db/sqlc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	dbStore, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET is not set")
	}

	clientID := os.Getenv("GCP_OAUTH_ID")
	if len(clientID) == 0 {
		log.Fatal("GCP_OAUTH_ID is not set")
	}

	clientSecret := os.Getenv("GCP_OAUTH_SECRET")
	if len(clientSecret) == 0 {
		log.Fatal("GCP_OAUTH_SECRET is not set")
	}

	redirectURLHost := os.Getenv("GCP_REDIRECT_URL_HOST")
	if len(redirectURLHost) == 0 {
		log.Fatal("REDIRECT_URL_HOST is not set")
	}

	// Set up the oauth config
	oauthConfig := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURLHost + "/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	grpcPort := getEnvOrDefault("GRPC_PORT", "8080")
	httpPort := getEnvOrDefault("PORT", "8081")

	// create server
	authenticationManager := auth.NewAuthenticationManager(jwtKey, oauthConfig)
	server := apipkg.NewServer(dbStore, authenticationManager)

	go runGrpcGatewayServer(httpPort)
	runGrpcServer(server, authenticationManager, grpcPort)
}

func runGrpcGatewayServer(httpPort string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// runtimeMux
	runtimeMux := runtime.NewServeMux(
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
			if customErr, ok := status.FromError(err); ok {
				http.Error(w, customErr.Message(), runtime.HTTPStatusFromCode(customErr.Code()))
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}),
	)

	grpcOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	endpoint := getEnvOrDefault("HOST", "localhost") + ":" + getEnvOrDefault("GRPC_PORT", "8080")

	// register
	pb.RegisterUserAuthenticationHandlerFromEndpoint(ctx, runtimeMux, endpoint, grpcOpts)
	pb.RegisterUserProfileHandlerFromEndpoint(ctx, runtimeMux, endpoint, grpcOpts)
	pb.RegisterExpenseManagementHandlerFromEndpoint(ctx, runtimeMux, endpoint, grpcOpts)

	// http server
	log.Printf("grpc-gateway server started on %s:%s", getEnvOrDefault("HOST", "localhost"), httpPort)
	err := http.ListenAndServe(getEnvOrDefault("HOST", "localhost")+":"+httpPort, runtimeMux)
	if err != nil {
		log.Fatalln(err)
	}
}

func accessibleRoles() map[string][]string {
	// Register each microservice with the roles that can access it.
	// The key is the microservice name and the value is the list of roles that can access it.
	// If the microservice is not registered, then it is accessible by anyone.
	const (
		userProfile       = "/apidefinitions.UserProfile/"
		expenseManagement = "/apidefinitions.ExpenseManagement/"
	)
	return map[string][]string{
		userProfile + "CreateUserProfileAPI":   {"admin", "user"},
		userProfile + "UpdateUserProfileAPI":   {"admin", "user"},
		userProfile + "GetUserProfileAPI":      {"admin", "user"},
		expenseManagement + "CreateExpenseAPI": {"admin", "user"},
		expenseManagement + "GetExpenseAPI":    {"admin", "user"},
		expenseManagement + "ListExpensesAPI":  {"admin", "user"},
		expenseManagement + "DeleteExpenseAPI": {"admin", "user"},
		expenseManagement + "UpdateExpenseAPI": {"admin", "user"},
	}
}

func runGrpcServer(server *apipkg.Server, authenticationManager *auth.AuthenticationManager, grpcPort string) {
	listener, err := net.Listen("tcp", getEnvOrDefault("HOST", "localhost")+":"+grpcPort)
	if err != nil {
		log.Fatalln(err)
	}

	// create auth interceptor
	authInterceptor := middlewares.NewAuthInterceptor(authenticationManager, accessibleRoles())

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middlewares.GrpcRequestValidator,
			authInterceptor.GrpcAuthInterceptor(),
		),
	)

	//register to GRPC server
	pb.RegisterUserAuthenticationServer(grpcServer, server)
	pb.RegisterUserProfileServer(grpcServer, server)
	pb.RegisterExpenseManagementServer(grpcServer, server)
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
