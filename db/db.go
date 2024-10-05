package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aveliap/transaction-go/config"
	_ "github.com/lib/pq"
)

func NewPostgresStorage(cfg config.Config) (*sql.DB, error)  {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping()
		err != nil{
			log.Fatal(err)
			return nil,err
		}

	return db,nil
}