package main

import (
	"flag"
	"server-a/config"
	"server-a/server"
)

var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*cfgPath)

	server.NewServer(cfg)
}
