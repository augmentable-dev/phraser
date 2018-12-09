package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/augmentable-opensource/phraser/pkg/api"
	"go.uber.org/zap"
)

func main() {
	ctx, closer := context.WithCancel(context.Background())

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := api.StartGRPC(ctx, ":50060", logger); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	log.Println("waiting for signal or errors")
	select {
	case sig := <-sigs:
		log.Printf("signal received, stopping signal=%s\n", sig.String())
	}

	closer()
}
