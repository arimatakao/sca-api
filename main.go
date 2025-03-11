package main

import (
	"flag"
	"log"

	"github.com/arimatakao/sca-api/config"
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
}
