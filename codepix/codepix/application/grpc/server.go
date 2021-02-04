package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/lucasres/code-pix/application/grpc/pb"
	"github.com/lucasres/code-pix/application/usecase"
	"github.com/lucasres/code-pix/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//StartGrpcServer cria o servidor grpc
func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixeKeyRepositoryDB{DB: database}
	pixUseCase := usecase.PixUseCase{Repository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixeServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start grpc server", err)
	}

	log.Printf("gRPC listiner at port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start grpc server", err)
	}
}
