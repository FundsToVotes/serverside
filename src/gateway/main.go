package main

import (
	"gateway/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//main is the main entry point for the server
func main() {
	//This code is untested!
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	//

	//Read the ADDR environment variable to get the address
	//	the server should listen on. If empty, default to ":80"
	addr := os.Getenv("GATEWAY_PORT")
	if len(addr) == 0 {
		addr = ":443"
	}

	//get the TLS key and cert paths from environment variables
	//this allows us to use a self-signed cert/key during development
	//and the Let's Encrypt cert/key in production
	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	if len(tlsCertPath) == 0 {
		log.Fatal("Error - could not find tlsCertPath enviroment variable")
	}

	if len(tlsKeyPath) == 0 {
		log.Fatal("Error - could not find tlsCertPath enviroment variable")
	}
	//Create a new mux for the web server.
	mux := http.NewServeMux()
	//Tell the mux to call your handlers.SummaryHandler function
	//	when the "/v1/summary" URL path is requested.
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/topten", handlers.DummyTopTenHandler)
	mux.HandleFunc("/bills", handlers.DummyBillsHandler)
	mux.HandleFunc("/billstest", handlers.BillsHandler)

	//Add the CORS middleware
	corsWrappedMux := handlers.NewCORSMiddleware(mux)

	//Start a web server listening on the address you read from
	//	the environment variable, using the mux you created as
	//	the root handler. Use log.Fatal() to report any errors
	//	that occur when trying to start the web server.
	log.Printf("the server is listening at port" + addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, corsWrappedMux)) //i dont know what this line does

}
