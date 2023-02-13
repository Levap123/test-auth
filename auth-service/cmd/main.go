package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Levap123/auth-service/internal/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := app.Run(); err != nil {
			log.Fatalln(err)
		}
	}()
	log.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
