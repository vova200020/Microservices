package main

import (
	"Auth/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models

}
func main(){
	log.Println("Starting Auth service")

	//TODO connect to db 


	//set up config
	app := Config{}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort)
		Handler.: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
		
	}

}
