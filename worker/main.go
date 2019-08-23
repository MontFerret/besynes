package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/MontFerret/besynes/worker/app"
	"os"
)

var (
	sub = flag.Int("sub", 5050, "ZeroMQ sub socket port")

	pub = flag.Int("pub", 5051, "ZeroMQ pub socket port")

	help = flag.Bool(
		"help",
		false,
		"show this list",
	)
)

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
		return
	}

	a, err := app.New(app.Settings{
		PubPort: *pub,
		SubPort: *sub,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	a.Run(context.Background())
}
