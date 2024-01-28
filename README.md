# Awesome Expense Tracker Backend

This is the backend service for the Awesome Expense Tracker application. It's written in Go and uses gRPC for communication.

---

## Directory Structure

- `api`: Contains the protobuf definitions and the generated Go code for the gRPC services.
- `pkg`: Contains the Go packages for the application.
  - `api`: Contains the implementation of the gRPC services.
  - `db`: Contains the database related code.

---

# What I learnt from this project?

## How to generate code from proto?

### What I used to do

I used to generate go code using the protobuf `protoc`. The protoc command was so tedious and unreliable for me. Providing the correct directory paths for all proto files and then some hacky ways to get protovalidator or grpc-gateway plugin or any other plugin to work. I am aware of creating a prototool.yaml to manage the proto plugins but it proved to be equally tedious to get all the plugins to work. (Maybe some learning is required here as well? 🙂)

### What I learnt

Using buf generate from [buf.build](https://buf.build/) is way easier to maintain and generate Go code from proto files. All the plugins and settings are in one place and it is easier to add or remove plugins this way.

## How to write SQL queries the better way?

### What I used to do

Even though I was used to just using GORM (which is not technically writing the SQL queries but rather an abstracted way of interacting with SQL Database). I had a goal of learning and practicing writing more SQL queries rather than depending on a library to do that for me.

### What I learnt

Even though my earlier approach looks good from learning perspective, this introduced other challenges of writing the corresponding go functions which parse the SQL queries and add the dynamic values (or as official website of SQLC describes it: _boilerplate SQL querying code_). I was not here to learn how to write more go functions but rather to only focus on SQL Queries. That was when I found out about SQLC. I did face challenges at first with the documentation being focused more for Postgresql rather than MySQL. Once I got the hang of it I was able to learn just enough for this project and write queries with SQLC generating boilerplate SQL querying Golang code for me.

> **NOTE**: Just as an added bonus I also refreshed my knowledge on database migrations and how to use them.

## Environment File

### What I used to do

This might be just a bad habit, but I used to write `export ENV_VAR=VALUE` in my makefile for local server startups. This was tedious and not very OS friendly way as windows uses `SET` to set env vars and linux uses `export`. (Yes I switch my dev environment sometimes so a consistent development experience is really important)

### What I learnt

Environment Files (.env). That's it. That's the solution and learning for me. It really helped me forget about exporting the variables again and again to really focus on the development without worrying about them.

## OAuth 2.0

### What I learnt

A lot was learnt in regards to OAuth 2.0, I even dared to create my own OAuth2.0 server implementation but decided against it (for now 😉). I had some experience in working with GCP's OAuth so I decided to integrate that first. I also wanted to provide users an ability to login/register using just the old fashioned email id and password and then later create their usernames. This led to me supporting both GCP's OAuth and also the JWT based authentication and then a dance of handling the different authentication methods started and I still don't believe that this is _THE_ perfect way to implement two authentications while maintaining fewer lines of codes but I have learnt a lot and came along way from what I used to do and I will continue to refactor or at least find ways to support multiple authentication methods.

## Interceptors and GRPC Gateway

### What I used to do

Having started my knowledge journey with MERN stack I was familiar with middlewares, that said the current implementation of interceptors in GRPC Gateway environment that I used to follow specially for authentication interceptors was somewhat in my opinion wrong. I used to implement the auth interceptors at GRPC gateway layer and reject a request if it ever fails the authentication or authorization barriers. However in my past projects (professional ones) we were not really utilizing the full potential of GRPC and GRPC Gateway but rather it was a REST based server only where GRPC was introduced at the start but was never developed upon later.

### What I learnt

Above mentioned approach of auth interceptors at the GRPC Gateway level was understandable in that environment as GRPC Ports are never utilized by the end customer, however with this project I wanted to support both RESTFul and GRPC calls both with GRPC Gateway and GRPC. However I recently revisited the actual GRPC Gateway github page: [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) and upon closer inspection I was able to deduce that this approach had a very big loophole for auth or any interceptors where the requests coming from the GRPC Gateway to my Server are only intercepted and processed whereas any request coming from GRPC directly will bypass any and all interceptors. This was really a big eye opener and I immediately moved all my interceptors to the GRPC level from GRPC Gateway. Now both the requests (RESTFul and GRPC) will have to pass through all the intended interceptors.

---

## Services

The backend provides the following gRPC services:

- `ExpenseManagement`: Provides methods for managing expenses. Implemented in [`expense_mgmt_api.go`](..\awesome-expense-tracker-backend\pkg\api\expense_mgmt_api.go)
- `UserAuth`: Provides methods for user authentication. Implemented in [`user_auth_api.go`](..\awesome-expense-tracker-backend\pkg\api\user_auth_api.go).
- `UserProfile`: Provides methods for managing user profiles. Implemented in [`user_profile_api.go`](..\awesome-expense-tracker-backend\pkg\api\user_profile_api.go).

---

## Environment Variables

The application uses the following environment variables:

- `APP_ENV`: The environment in which the application runs. Possible values are `local`, `dev`, `staging`, and `prod`. Default is `local`.
- `HOST`: The hostname where the application runs. Default is `localhost`.
- `PORT`: The port number where the HTTP server runs. Default is `8080`.
- `JWT_SECRET`: The secret key used for signing JSON Web Tokens (JWT) for user authentication and authorization.
- `GRPC_PORT`: The port number where the gRPC server runs. Default is `8081`.
- `MYSQLUSER`: The username for the MySQL database.
- `MYSQLPASSWORD`: The password for the MySQL database.
- `MYSQLADDR`: The address (hostname and port) of the MySQL server.
- `MYSQLDBNAME`: The name of the MySQL database.
- `GCP_REDIRECT_URL_HOST`: The hostname where the Google Cloud Platform (GCP) OAuth redirect URL is hosted.
- `GCP_OAUTH_SECRET`: The secret key for GCP OAuth authentication.
- `GCP_OAUTH_ID`: The client ID for GCP OAuth authentication.

You can set these environment variables in your shell, or you can use a `.env` file. If you're using a `.env` file, make sure to add it to your `.gitignore` file to prevent it from being committed to your repository.

---

## How to Run Locally

To run the backend service locally, you need to have Go installed. Then, you can run the following command:

```sh
go run server/main.go
```

---

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
