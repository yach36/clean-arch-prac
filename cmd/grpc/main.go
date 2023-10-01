package main

import (
	"database/sql"
	"fmt"
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
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPasswd := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPasswd, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	serverAddr := viper.GetString("grpc.server.address")
	ln, err := net.Listen("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	deliveryGrpcHandler.NewUserServerGrpc(server, userUsecase)
	if err := server.Serve(ln); err != nil {
		panic(err)
	}
	log.Printf("grpc server start running on %s\n", serverAddr)
}
