package handlers

import (
	"encoding/json"
	"fmt"
	"gateway/useDB_for_handler"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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

	//If Database doesn't have data, fetch it from Opensecrets, then cache it in the database
	if toptenToReturn.Attributes.CandName == "" {
		log.Printf("Fetching new data from Opensecrets then caching it")

		//API Key
		openSecretsAPIKey := os.Getenv("SERVERSIDE_OPENSECRETS_API_KEY")

		if len(openSecretsAPIKey) == 0 {
			log.Fatal("Error - could not find SERVERSIDE_OPENSECRETS_API_KEY enviroment variable")
		}

		client := &http.Client{}
		//Make the request
		request_url := "http://www.opensecrets.org/api/?method=candIndustry&cid=" + crp_id + "&cycle=2020&apikey=" + openSecretsAPIKey + "&output=json"
		req, err := http.NewRequest("GET", request_url, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Set("User-Agent", "Golang_Funds_To_Votes_Bot")
		response, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		//Read the Response
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Error in reading response")
			log.Fatal(err)
		}

		//Create an empty instance of a Congressperson to recieve JSON
		congressperson := &useDB_for_handler.Congressperson{}
		json.Unmarshal(responseData, congressperson)

		//Inserting into the db
		err = useDB_for_handler.Insert(db, *congressperson)
		if err != nil {
			if strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
				log.Printf("Entry for member " + congressperson.Response.Industries.Attributes.CandName + " already exists, updating instead")
				//UPDATE FUNCTION TO DO for when data is out of date
			} else {
				log.Printf("Insert member "+congressperson.Response.Industries.Attributes.CandName+" failed with error %s", err)
				//Will just return the empty struct in this case
			}
		}

		//Repopulate the return struct with the new values
		//Select from DB
		toptenToReturn, err = useDB_for_handler.Select(db, crp_id)
		if err != nil {
			log.Printf("Main: Error in Selecting "+crp_id+", failed with error %s", err)
			return
		}
	}

	//Finally, respond with a JSON-encoded version of the toptenToReturn struct.
	enc := json.NewEncoder(w)
	if err := enc.Encode(toptenToReturn); err != nil {
		fmt.Printf("error encoding struct into JSON: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	//TODO - SOMETHIGN ABOUT CLOSING THE HTML STREAM TO AVOID WASTING RESOURCES

}
