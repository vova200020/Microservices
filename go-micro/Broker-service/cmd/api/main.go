package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker sservice on port %s/n", webPort)

	//define htp server
	srv := &http.Server{
		Addr:    fmt.Sprint(":%s", webPort),
		Handler: app.routes(),
	}
	//start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
		log.Default()
	}
}
