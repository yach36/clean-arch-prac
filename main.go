package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/spf13/viper"
	adapterHttp "github.com/yach36/clean-arch-prac/adapter/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

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

	r := adapterHttp.NewRouter()

	serverAddr := viper.GetString("server.address")
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
