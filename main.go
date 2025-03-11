package main

import (
	"flag"
	"log"

	"github.com/arimatakao/sca-api/config"
	"github.com/arimatakao/sca-api/server"
)

var cfgPath *string

func init() {
	cfgPath = flag.String("config", "./config.yaml", "config path")
	flag.Parse()
}

func main() {
	if err := config.LoadConfig(*cfgPath); err != nil {
		log.Fatal(err)
	}
	if err := server.New().Run(); err != nil {
		log.Fatal(err)
	}
}
