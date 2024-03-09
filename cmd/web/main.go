package main

import (
	"fmt"
	"github.com/fouched/go-sample-web/internal/config"
	"log"
	"net/http"
)

const port = ":8000"

var app config.AppConfig

func main() {

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}
	fmt.Println(fmt.Sprintf("Starting application on %s", port))

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
