package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nollidnosnhoj/simpbb/internal/server"
)

func main() {
	cctx, cancel := context.WithCancel(context.Background())

	// start server gorountine
	go server.Start(cctx)

	// wait for signal to initiate shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	sig := <-quit

	log.Println("Shutting down...")
	cancel()

	if sig == syscall.SIGUSR1 {
		os.Exit(1)
	}
}