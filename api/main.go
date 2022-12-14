package main

import (
	"fmt"
	"log"
	"net/http"

	"go-devbook-api/src/config"
	"go-devbook-api/src/router"
)

func main() {
	config := config.Load()

	r := router.Build()

	port := fmt.Sprintf(":%d", config.Port)

	uri := config.UriConection
	log.Println(fmt.Sprintf("db connection url :%s", uri))

	log.Println(fmt.Sprintf("Starting in localhost:%s", port))
	log.Fatal(http.ListenAndServe(port, r))
}
