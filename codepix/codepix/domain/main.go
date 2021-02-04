package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lucasres/code-pix/application/grpc"
	"github.com/lucasres/code-pix/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.CreateDB(os.Getenv("ENV"))
	grpc.StartGrpcServer(database, 50051)
}
