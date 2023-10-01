package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	deliveryHttp "github.com/yach36/clean-arch-prac/delivery/http"
	"github.com/yach36/clean-arch-prac/delivery/http/controller"
	"github.com/yach36/clean-arch-prac/infra/postgres"
	"github.com/yach36/clean-arch-prac/usecase"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

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

	userRepo:= postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	r := deliveryHttp.NewRouter(userController)

	serverAddr := viper.GetString("api.server.address")
	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}
	timeout, _ := strconv.Atoi(viper.GetString("context.timeout"))
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
	log.Printf("server start running on %s\n", serverAddr)
}
