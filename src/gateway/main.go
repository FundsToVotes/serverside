package main

import (
	"gateway/handlers"
	"log"
	"net/http"
	"os"
)

//main is the main entry point for the server
func main() {
	//This code is untested!
	// loads values from .env into the system
	/*	if err := godotenv.Load("gatewayVariables.env"); err != nil {
		//read working directory files
		/*
			files, err := ioutil.ReadDir("./")
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range files {
				fmt.Println(f.Name())
			} */
	//	log.Printf(err.Error())
	//		log.Fatal("No .env file found")
	//	} //Replaced with putting the env files in bash scripts for building the docker containers
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

	//THE FOLLOWING CODE IS UNTESTED
	propublicaCongressAPIKey := os.Getenv("SERVERSIDE_APP_PROPUBLICA_CONGRESS_API_KEY")

	if len(propublicaCongressAPIKey) == 0 {
		log.Fatal("In main, Error - could not find SERVERSIDE_APP_PROPUBLICA_CONGRESS_API_KEY enviroment variable")
	}

	//Create a new mux for the web server.
	mux := http.NewServeMux()
	//Tell the mux to call your handlers.SummaryHandler function
	//	when the "/v1/summary" URL path is requested.
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/topten", handlers.DummyTopTenHandler)
	mux.HandleFunc("/billsdummy", handlers.DummyBillsHandler)
	mux.HandleFunc("/bills", handlers.BillsHandler)

	//Add the CORS middleware
	corsWrappedMux := handlers.NewCORSMiddleware(mux)

	//Start a web server listening on the address you read from
	//	the environment variable, using the mux you created as
	//	the root handler. Use log.Fatal() to report any errors
	//	that occur when trying to start the web server.
	log.Printf("the server is listening at port" + addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, corsWrappedMux)) //i dont know what this line does

}
