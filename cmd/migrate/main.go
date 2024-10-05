package main

import (
	"log"
	"os"

	"github.com/aveliap/transaction-go/config"
	"github.com/aveliap/transaction-go/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main(){
	cfg:= config.Envs
	db, err := db.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migration",
		"postgres",
		driver)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		if err := m.Up(); 
		err!=nil && err != migrate.ErrNoChange{
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Successfully migrates")	
	}

	if cmd == "down" {
		if err := m.Down(); 
		err!=nil && err != migrate.ErrNoChange{
			log.Fatalf("Failed to revert migrations: %v", err)
		}	
	log.Println("Successfully revert")
	}
}