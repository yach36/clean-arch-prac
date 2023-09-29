package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	deliveryHttp "github.com/yach36/clean-arch-prac/delivery/http"
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

	r := deliveryHttp.NewRouter()

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

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("server start running on %s\n", serverAddr)
}
