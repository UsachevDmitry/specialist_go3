package main

import (
	"BankstoreOpk/api"
	db "BankstoreOpk/db/sqlc"
	"BankstoreOpk/utils"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://postgres:P@ssw0rd@localhost:5432/bankdb?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)



func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("can not config file", err)
	}
	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}
	defer pool.Close()

	store := db.NewStore(pool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}
}