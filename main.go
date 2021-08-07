package main

import (
	"log"
	"net/http"
)

const httpPort = ":8081"

func main() {
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
