package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/machilan1/go_prc/postgres"
)

type application struct {
	logger *slog.Logger
	store  *postgres.Store
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Can't load dotenv")
	}

	dataSourceName := os.Getenv("DB_URL")
	addr := flag.String("addr", ":4000", "HTTP network port number")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	store, err := postgres.NewStore(dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("Database not connected")
	}

	app := application{
		logger: logger,
		store:  store,
	}

	logger.Info("Starting server at ", slog.String("addr", *addr))
	err = http.ListenAndServe(*addr, app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
