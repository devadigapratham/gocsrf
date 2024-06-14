package main

import (
	"log"

	"github.com/devadigapratham/gocsrf/db"
	"github.com/devadigapratham/gocsrf/server"
	"github.com/devadigapratham/gocsrf/server/middleware/myJwt"
)

var host = "localhost"
var port = "9000"

func main() {

	db.InitDB()

	jwtErr := myJwt.InitJWT()
	if jwtErr != nil {
		log.Println("Error initializing the JWT!")
		log.Fatal(jwtErr)
	}

	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting the server")
		log.Fatal(serverErr)
	}
}
