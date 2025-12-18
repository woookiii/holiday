package main

import (
	"flag"
	"server-a/config"
	"server-a/server"

	"go.uber.org/fx"
)

var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*cfgPath)

	fx.New(
		fx.Provide(func() *config.Config { return cfg }),

		fx.Provide(server.NewServer),

		fx.Invoke(func(_ *server.Server) {}),
	).Run()

}
