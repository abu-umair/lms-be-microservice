package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/abu-umair/lms-be-microservice/internal/handler"
	"github.com/abu-umair/lms-be-microservice/internal/repository"
	"github.com/abu-umair/lms-be-microservice/internal/service"
	"github.com/abu-umair/lms-be-microservice/pb/auth"
	"github.com/abu-umair/lms-be-microservice/pkg/database"
	"github.com/abu-umair/lms-be-microservice/pkg/grpcmiddleware"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("Error when listening %v", err)
	}

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))

	log.Println("Connected to DB")
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
		),
	)

	auth.RegisterAuthServiceServer(serv, authHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Printf("Reflection is Registered.")

	}

	log.Println("Server is running on :50051 port")
	if err := serv.Serve(lis); err != nil {
		log.Panicf("Server is error %v", err)
	}
}
