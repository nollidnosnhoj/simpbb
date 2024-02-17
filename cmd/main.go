package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const name = "simpbb"
const about = "simpbb - a simple forum"
const description = "A simple forum written in Go"

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = about
	app.Description = description
	app.Commands = []*cli.Command{
		StartCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}