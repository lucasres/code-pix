package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

//StartGrpcServer cria o servidor grpc
func StartGrpcServer(db *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

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
