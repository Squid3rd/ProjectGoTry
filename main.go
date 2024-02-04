package main

import (
	"context"
	"log"
	"os"

	"github.com/Squid3rd/FirstGoProject/config"
	"github.com/Squid3rd/FirstGoProject/pkg/database"
)

func main() {
	ctx := context.Background()

	// Initial Config
	cfg := config.LoadConfig(func() string {

		if len(os.Args) < 2 {
			log.Fatal("Please provide .env Required")
		}

		return os.Args[1]
	}())

	// Connect to DB
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	log.Println(db)

	log.Println(cfg)
}
