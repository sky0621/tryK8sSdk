package main

import (
	"github.com/sky0621/tryK8sSdk"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", tryK8sSdk.HelloHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
