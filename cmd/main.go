package main

import (
	"log"

	"github.com/aveliap/transaction-go/cmd/api"
	"github.com/aveliap/transaction-go/config"
	"github.com/aveliap/transaction-go/db"
)

func main(){
	cfg:= config.Envs
	db, err := db.NewPostgresStorage(cfg)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully Connect")

	server:=api.NewAPIServer(":8080", db)
	if err:= server.Run()
		err != nil{
		log.Fatal(err)
		}
}