package main

import (
	"flag"
	"server-a/config"
	"server-a/src"
)

var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*cfgPath)

	src.NewServer(cfg)
}
