package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//BillsHandler returns Alma Adams's JSON data, and nobody else's. It does check that a name was provided
func BillsHandler(w http.ResponseWriter, r *http.Request) {
	//Add an HTTP headers to the response with the name
	w.Header().Set("Content-Type", "application/json")

	//Get the `url` query string parameter value from the request.
	keys, ok := r.URL.Query()["name"]

	//If not supplied, respond with an http.StatusBadRequest error.
	if !ok || len(keys[0]) < 1 || len(keys) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "StatusBadRequestError: Please provide a representative name", 400)
		return
	}

	key := keys[0]
	fmt.Println("Request made for Representative " + key)

	// SOME SORTA CODE HERE TO GO FROM NAME TO CRP ID, WHICH I MIGHT CO-DEVELOP WITH THE TOPTEN
	//for now, the next line is temporary
	member_id := "K000388"

	//Fetch all bills a member of congress has recently voted on
	membersAPIResponse, err := fetch_member_bills(member_id)
	if err != nil {
		log.Fatal(err)
	}

	billsToServe := &BillsToServe{}
	billsToServe.Status = membersAPIResponse.Status
	//billsToServe.Results[0].

	//For each bill, fetch the bill details
	for _, eachBill := range membersAPIResponse.Results[0].Votes {
		bill_slug := strings.TrimSuffix(eachBill.Bill.BillID, "-117")
		//TODO - PUT THE "IF LATEST ACTION IS A DESRIRABLE ONE" STATEMENT HERE
		billsAPIResponse, err := fetch_bill_details(bill_slug)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Primary subject: " + billsAPIResponse.Results[0].PrimarySubject)
	}
}

//This function fetches the bills the member of congress has recently voted on
func fetch_member_bills(member_id string) (*MembersAPIResponse, error) {
	//Create the HTTP Request
	client := &http.Client{}
	request_url := "https://api.propublica.org/congress/v1/members/" + member_id + "/votes.json"
	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		return nil, err
	}

	//Set the Request headers
	req.Header.Set("User-Agent", "Golang_Funds_To_Votes_Bot")
	req.Header.Set("X-API-Key", "bUEZbt82MwpoNooSEbGjITvWKC703nY2isRQVa5Y")

	//Make the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//Read the request into the responseData variable
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Fetched bills recently voted on for" + member_id + "without errors")

	//Read the fetched data into the MemberAPIResponse struct
	membersAPIResponse := &MembersAPIResponse{}
	json.Unmarshal(responseData, membersAPIResponse)

	return membersAPIResponse, nil
}

func fetch_bill_details(bill_slug string) (*BillsAPIResponse, error) {
	//Create the Http Request
	client := &http.Client{}
	request_url := "https://api.propublica.org/congress/v1/117/bills/" + bill_slug + ".json"
	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		return nil, err
	}

	//Set the Request Headers
	req.Header.Set("User-Agent", "Golang_Funds_To_Votes_Bot")
	req.Header.Set("X-API-Key", "bUEZbt82MwpoNooSEbGjITvWKC703nY2isRQVa5Y")

	//Make the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//Read the request into the responseData variable
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//Create an empty instance of a BillsAPIResponse to recieve JSON
	billsAPIResponse := &BillsAPIResponse{}
	json.Unmarshal(responseData, billsAPIResponse)

	if len(billsAPIResponse.Results) > 0 {
		//FIX THIS and MAKE THIS BETTER
		fmt.Println("Primary subject: " + billsAPIResponse.Results[0].PrimarySubject)
		fmt.Println("Short title: " + billsAPIResponse.Results[0].ShortTitle)
		fmt.Println()
	} else {
		fmt.Println("No data returned for " + bill_slug)
	}

	return billsAPIResponse, nil
}
