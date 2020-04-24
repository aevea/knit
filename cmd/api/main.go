package main

import (
	"log"
	"net/http"

	"github.com/aevea/merge-master/api"
	"github.com/aevea/merge-master/api/generated"
	"github.com/pacedotdev/oto/otohttp"
)

func main() {
	g := api.HealthcheckService{}
	server := otohttp.NewServer()
	generated.RegisterHealthcheck(server, g)
	http.Handle("/v1/", server)
	log.Println("Server starting on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
