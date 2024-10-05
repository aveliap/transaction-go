package main

import (
	"log"

	"github.com/aveliap/transaction-go/cmd/api"
	"github.com/aveliap/transaction-go/db"
)

func main(){

	db, err := db.NewPostgresStorage(postgres.Config{
		User: "postgres",
		Password: "postgres",
		Addr: "localhost",
		Port: "5432",
		DBName: "transaction_go",

	})
	server:=api.NewAPIServer(":8080", nil)
	if err:= server.Run()
		err != nil{
		log.Fatal(err)
		}
}