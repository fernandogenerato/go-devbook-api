package main

import (
	"fmt"
	"log"
	"net/http"

	"go-devbook-api/src/config"
	"go-devbook-api/src/router"
)

// func init() {
// 	key := make([]byte, 64)
// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	encoding := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(encoding)
// }

func main() {
	config := config.Load()

	r := router.Build()

	port := fmt.Sprintf(":%d", config.Port)

	uri := config.UriConection
	log.Println(fmt.Sprintf("db connection url :%s", uri))

	log.Println(fmt.Sprintf("Starting in localhost:%s", port))
	log.Fatal(http.ListenAndServe(port, r))
}
