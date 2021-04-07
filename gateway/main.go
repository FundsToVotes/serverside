package main

import (
	"assignments-jhoupps/servers/gateway/handlers"
	"log"
	"net/http"
	"os"
)

//main is the main entry point for the server
func main() {
	//Read the ADDR environment variable to get the address
	//	the server should listen on. If empty, default to ":80"
	addr := os.Getenv("ADDR")
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

	//Start a web server listening on the address you read from
	//	the environment variable, using the mux you created as
	//	the root handler. Use log.Fatal() to report any errors
	//	that occur when trying to start the web server.
	log.Printf("the server is listening at port" + addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux)) //i dont know what this line does

}