package main

import (
	"src/internal/di"

	"go.uber.org/fx"
)

// точка входа, здесь вызовем di и запустим игру

func main() {
	fx.New(di.CreateApp()).Run()
}
