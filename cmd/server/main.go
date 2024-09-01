package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/prodsub/pb/gen"
	"github.com/prodsub/pkg"
	"github.com/prodsub/pkg/db"
	"github.com/prodsub/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	log.Println("Starting ProductSubscription Service")
	gdb := initDb()

	conf := pkg.Config{
		Port: 50051,
	}
	runGrpcServer(gdb, conf)

	log.Println("Exiting ProductSubscription Service")
}

func initDb() *gorm.DB {
	GormDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres dbname=prodsub_db port=5432 sslmode=disable",
	}))

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %s", err.Error())
	}

	// Migrate the schema (create the Servvice tables)
	if err := GormDb.AutoMigrate(db.Product{}, db.Subscription{}); err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}

	return GormDb
}

func runGrpcServer(gdb *gorm.DB, conf pkg.Config) {

	serv := service.NewProductSubscriptionService(db.NewProductRepo(gdb), db.NewSubscriptionRepo(gdb))

	// Set up gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, serv)
	pb.RegisterSubscriptionServiceServer(grpcServer, serv)

	// support reflection
	reflection.Register(grpcServer)

	fmt.Printf("gRPC server is running on port %d...", conf.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
