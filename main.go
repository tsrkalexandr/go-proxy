package main

import (
	"flag"
	"fmt"

	"github.com/tsrkalexandr/go-proxy/config"
	"github.com/tsrkalexandr/go-proxy/proxy"
	"github.com/tsrkalexandr/go-proxy/storage"
)

func main() {
	var confFile string

	// get config file path from input argument
	flag.StringVar(&confFile, "config", "./etc/config.yml", "Path to configuration file")
	flag.Parse()

	// get config file path from input argument
	cfg, err := config.NewFromFile(confFile)
	if err != nil {
		fmt.Printf("failed to read config:\n%s\nUse default configuration\n", err)
	}

	server := proxy.NewServer(cfg, storage.NewStorage())

	if err := server.Start(); err != nil {
		fmt.Println(err)
	}
}
