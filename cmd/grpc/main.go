package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	deliveryGrpcHandler "github.com/yach36/clean-arch-prac/delivery/grpc/handler"
	"github.com/yach36/clean-arch-prac/infra/postgres"
	"github.com/yach36/clean-arch-prac/usecase"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	dbConn := postgres.NewPostgresConnector()
	userRepo := postgres.NewUserRepository(dbConn.DB)
	userUsecase := usecase.NewUserUsecase(userRepo)

	serverAddr := viper.GetString("grpc.server.address")
	ln, err := net.Listen("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	deliveryGrpcHandler.NewUserServerGrpc(server, userUsecase)
	if err := server.Serve(ln); err != nil {
		log.Fatalln(err)
	}
	log.Printf("grpc server start running on %s\n", serverAddr)
}
