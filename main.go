package main

import (
	"api/api"
	db "api/db/sqlc"
	"api/util"
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	// Load config from .env file
	config, err := util.Loadconfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Connect to DB using config values from .env file
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	fmt.Println(dbSource)
	conn, err := sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// run db migrations
	runDBMigration(config.MigrationURL, dbSource)

	// Create new store and server
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migration instance: ", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migration up: ", err)
	}

	log.Println("db migrated successfully")
}
