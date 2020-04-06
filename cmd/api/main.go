package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/commitsar-app/merge-master/api/rest/v1"
)

func main() {
	restServer := rest.NewServer("merge-master")

	address := "localhost"
	port := "3000"
	serveAddr := net.JoinHostPort(address, port)
	server := &http.Server{
		Handler:      restServer.Router,
		Addr:         serveAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
