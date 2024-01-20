# Awesome Expense Tracker Backend

This is the backend service for the Awesome Expense Tracker application. It's written in Go and uses gRPC for communication.

## Directory Structure

- `api`: Contains the protobuf definitions and the generated Go code for the gRPC services.
- `pkg`: Contains the Go packages for the application.
  - `api`: Contains the implementation of the gRPC services.
  - `db`: Contains the database related code.

## Services

The backend provides the following gRPC services:

- `ExpenseManagement`: Provides methods for managing expenses. Implemented in [`api_grpc.pb.go`](..\awesome-expense-tracker-backend\api\api_grpc.pb.go).
- `UserAuth`: Provides methods for user authentication. Implemented in [`user_auth_api.go`](..\awesome-expense-tracker-backend\pkg\api\userauth\user_auth_api.go).

## Environment Variables

The application uses the following environment variables:

- `HOST`: The hostname where the application runs. Default is `localhost`.
- `PORT`: The port number where the HTTP server runs. Default is `8080`.
- `JWT_SECRET`: The secret key used for signing JSON Web Tokens (JWT) for user authentication and authorization.
- `GRPC_PORT`: The port number where the gRPC server runs. Default is `8081`.
- `MYSQLUSER`: The username for the MySQL database.
- `MYSQLPASSWORD`: The password for the MySQL database.
- `MYSQLADDR`: The address (hostname and port) of the MySQL server.
- `MYSQLDBNAME`: The name of the MySQL database.

You can set these environment variables in your shell, or you can use a `.env` file. If you're using a `.env` file, make sure to add it to your `.gitignore` file to prevent it from being committed to your repository.

## How to Run Locally

To run the backend service locally, you need to have Go installed. Then, you can run the following command:

```sh
go run server/main.go
```

## How to Deploy to the Cloud

To deploy the backend service to a cloud provider, you will need to containerize the application using Docker and then deploy the Docker container to your cloud provider. Here are the general steps:

1. Build the Docker image:

```sh
docker build -t awesome-expense-tracker-backend .
```

2. Push the Docker image to a registry:

```sh
docker push awesome-expense-tracker-backend
```

3. Deploy the Docker image to your cloud provider. The exact command will depend on your cloud provider. For example, on Google Cloud Platform, you would use:

```sh
gcloud run deploy --image gcr.io/your-project-id/awesome-expense-tracker-backend
```

Please refer to your cloud provider's documentation for the exact steps
