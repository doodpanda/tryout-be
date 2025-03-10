package main

import (
	"context"
	"log"
	"time"

	"github.com/doodpanda/tryout-backend/internal/config"
	"github.com/doodpanda/tryout-backend/platform/database"
)

// @title API Service for Tryout
// @version 1.0
// @description API Service for Tryout
// @termsOfService
// @contact.name API Maintainer
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host null
// @BasePath /
func main() {
	dbtype := config.GetEnv("DATABASE_TYPE", "pgx")
	ctx := context.Background()

	if dbtype == "pgx" {
		database.ConnectPostgres(ctx)
	} else {
		panic("invalid database type")
	}

	prod := false

	app := NewApp(&appConfig{
		prod: prod,
		db:   database.DB,
	})

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			database.DB.Close()
			return nil
		},
		"http-server": func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	port := ":" + config.GetEnv("SERVER_PORT", "8080")

	if err := app.Listen(port); err != nil {
		log.Panic(err)
	}

	<-wait
}
