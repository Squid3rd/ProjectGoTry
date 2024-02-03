package main

import (
	"context"
	"log"
	"os"

	"github.com/Squid3rd/FirstGoProject/config"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initial Config
	cfg := config.LoadConfig(func() string {

		if len(os.Args) < 2 {
			log.Fatal("Please provide .env Required")
		}

		return os.Args[1]
	}())

	log.Println(cfg)
}
