package main

import (
	"log"
	"net"
	"os"

	"github.com/abu-umair/lms-be-microservice/internal/handler"
	"github.com/abu-umair/lms-be-microservice/pb/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	godotenv.Load()
	
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("Error when listening %v", err)
	}

	serviceHandler := handler.NewServiceHandler()

	serv := grpc.NewServer()

	service.RegisterHelloWorldServiceServer(serv, serviceHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Printf("Reflection is Registered.")

	}

	log.Println("Server is running on :50051 port")
	if err := serv.Serve(lis); err != nil {
		log.Panicf("Server is error %v", err)
	}
}
