package dpi

import (
	"context"
	"fmt"
	"log"

	infrastructures "challenge/internal/infrastructure"

	"github.com/jackc/pgx/v5/pgxpool"
)

// simple singleton dependency injection because limitation of resources
var PGXPool *pgxpool.Pool // use connection pool because its concurent safe

func SetupDPI(env *infrastructures.Env) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	dbpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}
	PGXPool = dbpool
}
