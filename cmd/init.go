package main

import (
	"github.com/joho/godotenv"

	"github.com/phuwn/tools/log"
	"github.com/phuwn/tools/util"

	"github.com/phuwn/yahee/src/server"
	"github.com/phuwn/yahee/src/store"
)

// init server stuff
func init() {
	env := util.Getenv("RUN_MODE", "")
	if env == "local" || env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	store := store.New()
	server.NewServerCfg(store)
}