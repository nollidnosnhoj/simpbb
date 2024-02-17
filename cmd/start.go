package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nollidnosnhoj/simpbb/internal/server"
	"github.com/urfave/cli/v2"
)

var StartCommand = &cli.Command{
	Name: "start",
	Usage: "Start the server",
	Action: startServer,
}

func startServer(c *cli.Context) error {
	cctx, cancel := context.WithCancel(context.Background())

	s := server.NewServer()

	// start server gorountine
	go s.Start(cctx)

	// wait for signal to initiate shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	sig := <-quit

	log.Println("Shutting down...")
	cancel()

	if sig == syscall.SIGUSR1 {
		os.Exit(1)
		return nil
	}

	return nil
}