package handlers

import (
	"api_call/useDB"
	"fmt"
	"log"
	"net/http"
)

//HelloHandler is a very basic handler for the purpose of getting a bare bones server up and running
func NotYetWorkingTopTenHandler(w http.ResponseWriter, r *http.Request) {

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
	db, err := useDB.Connect()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")
	err = useDB.CreateTopTenTable(db)
	if err != nil {
		log.Printf("Create top ten table failed with error %s", err)
		return
	}

	//Retrieve data from the internal database

	//Select from DB
	toptenToReturn, err := useDB.Select(db, crp_id)
	if err != nil {
		log.Printf("Main: Error in Selecting "+crp_id+", failed with error %s", err)
		return
	}

	fmt.Println("name to return: " + toptenToReturn.Attributes.CandName)
	fmt.Println("Industry1 to return: " + toptenToReturn.Industry[0].IndustryName)

}
