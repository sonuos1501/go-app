package main

import (
	"github.com/sonuos1501/go-app/worker/boostrap"
	"go.uber.org/fx"
)

func main() {

	fx.New(fx.Options(boostrap.All()...)).Run()
}
