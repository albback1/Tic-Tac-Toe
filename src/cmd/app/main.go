package main

import (
	"tictac/internal/di"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(di.CreateApp())
	app.Run()
}
