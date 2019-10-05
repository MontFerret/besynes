package main

import "github.com/MontFerret/besynes/internal"

func main() {
	app, err := internal.New()

	if err != nil {
		panic(err)
	}

	err = app.Run()

	if err != nil {
		panic(err)
	}
}
