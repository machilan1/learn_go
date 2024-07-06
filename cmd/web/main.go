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

	// 載入環境變數
	err := godotenv.Load()
	if err != nil {
		panic("Can't load dotenv")
	}
	dataSourceName := os.Getenv("DB_URL")

	// 載入flags

	addr := flag.String("addr", ":4000", "HTTP network port number")
	flag.Parse()

	// Logger 初始化
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Database Client 初始化
	store, err := postgres.NewStore(dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("Database not connected")
	}

	// 應用程式初始化
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
