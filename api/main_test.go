package api

import (
	"context"
	"fmt"
	"testing"
	"time"

	"log"
	"os"

	db "exchange-rate-api/db/postgres"
	"exchange-rate-api/pkg"
	"exchange-rate-api/tools"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	server *apiServer
)

func TestMain(m *testing.M) {
	// Scrapping and save into JSON file
	pkg.Scrapping()

	// Load Configuration
	configAPI, err := tools.LoadConfig("..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	//--- Parseconfig ---//
	// "postgresql://pg:secret@localhost:5432/exchange?sslmode=disable"
	configPG, err := pgxpool.ParseConfig(configAPI.DBAddress) // Using environment variables instead of a connection string.
	if err != nil {
		log.Fatal(err)
	}

	//--- pgxpool ---//
	dbpool, err := pgxpool.New(context.Background(), configPG.ConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	pgPool := db.NewDB(dbpool)

	// create context for timeout duration
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	// insert data to postgresql
	err = pgPool.SaveDataToDB(ctx)
	if err != nil {
		log.Println(err)
	}

	server = NewApiServer(ctx, pgPool)
	server.StartServer(configAPI.ServerAddress)

	os.Exit(m.Run())
}
