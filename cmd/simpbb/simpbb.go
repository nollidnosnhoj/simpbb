package main

import (
	"log"
	"os"

	"github.com/nollidnosnhoj/simpbb/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Commands = commands.Commands

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}