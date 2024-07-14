package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/pleimer/ticketer/server/cmd"
)

type Server struct {
	cmd.Start   `command:"start"`
	cmd.Migrate `command:"migrate"`
}

func main() {
	app := Server{}

	parser := flags.NewParser(&app, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		panic(err)
	}
}
