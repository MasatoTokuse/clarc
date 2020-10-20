package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() error {

	router, err := getRouter()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    ":8800",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-signalChannel

	log.Println("Graceful Shutdown...")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		log.Println(err.Error())
	}

	return nil
}
