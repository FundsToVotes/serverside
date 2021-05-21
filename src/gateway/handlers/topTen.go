package handlers

import (
	"encoding/json"
	"fmt"
	"gateway/useDB_for_handler"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//HelloHandler is a very basic handler for the purpose of getting a bare bones server up and running
func TopTenHandler(w http.ResponseWriter, r *http.Request) {

	//Add an HTTP headers to the response with the name
	w.Header().Set("Content-Type", "application/json")

	//Get the `url` query string parameter value from the request.
	keys, ok := r.URL.Query()["crp_id"]

	//If not supplied, respond with an http.StatusBadRequest error.
	if !ok || len(keys[0]) < 1 || len(keys) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "StatusBadRequestError: Please provide a crp_id", 400)
		return
	}

	crp_id := keys[0]
	log.Printf("Request made for Representative " + crp_id)

	//Connect to the DB
	//Todo - decide if this should go in main, not the handler code

	//*********Managing the DB*******
	db, err := useDB_for_handler.Connect()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")
	err = useDB_for_handler.CreateTopTenTable(db)
	if err != nil {
		log.Printf("Create top ten table failed with error %s", err)
		return
	}

	//Retrieve data from the internal database

	//Select from DB
	toptenToReturn, err := useDB_for_handler.Select(db, crp_id)
	if err != nil {
		log.Printf("Main: Error in Selecting "+crp_id+", failed with error %s", err)
		return
	}

	//Finally, respond with a JSON-encoded version of the toptenToReturn struct.
	enc := json.NewEncoder(w)
	if err := enc.Encode(toptenToReturn); err != nil {
		fmt.Printf("error encoding struct into JSON: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	//TODO - SOMETHIGN ABOUT CLOSING THE HTML STREAM TO AVOID WASTING RESOURCES

}
