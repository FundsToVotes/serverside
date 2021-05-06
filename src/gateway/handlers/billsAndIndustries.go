package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//BillsHandler returns Alma Adams's JSON data, and nobody else's. It does check that a name was provided
func BillsHandler(w http.ResponseWriter, r *http.Request) {
	//Add an HTTP headers to the response with the name
	w.Header().Set("Content-Type", "application/json")

	//Get the `url` query string parameter value from the request.
	keys, ok := r.URL.Query()["member_id"]

	//If not supplied, respond with an http.StatusBadRequest error.
	if !ok || len(keys[0]) < 1 || len(keys) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "StatusBadRequestError: Please provide a member_id", 400)
		return
	}

	member_id := keys[0]
	fmt.Println("Request made for Representative " + member_id)

	//member_id := "K000388" //For Debugging

	//Fetch all bills a member of congress has recently voted on
	membersAPIResponse, err := fetch_member_bills(member_id)
	if err != nil {
		log.Fatal(err)
	}

	billsToServe := &BillsToServe{}
	billsToServe.Status = membersAPIResponse.Status
	billsToServe.MemberID = membersAPIResponse.Results[0].MemberID
	billsToServe.TotalVotes = membersAPIResponse.Results[0].TotalVotes
	billsToServe.Offset = membersAPIResponse.Results[0].Offset

	//For each bill, fetch the bill details
	for _, eachMembersBill := range membersAPIResponse.Results[0].Votes {
		bill_slug := strings.TrimSuffix(eachMembersBill.Bill.BillID, "-117")

		//Check if in the list of bill Latest Actions we want
		if true /*isActuallyABill(eachMembersBill.Question, eachMembersBill.Bill.BillID)*/ {

			//TODO - PUT THE "IF LATEST ACTION IS A DESRIRABLE ONE" STATEMENT HERE
			//fmt.Println(eachMembersBill)
			billsAPIResponse, err := fetch_bill_details(bill_slug)
			if err != nil {
				log.Fatal(err)
			}

			if len(billsAPIResponse.Results) > 0 {

				//From Members API
				newVotesStruct := VotesToServe{}
				newVotesStruct.Session = eachMembersBill.Session
				newVotesStruct.RollCall = eachMembersBill.RollCall
				newVotesStruct.VoteURI = eachMembersBill.VoteURI
				newVotesStruct.Position = eachMembersBill.Position

				//From Bills API, and Members API
				newVotesStruct.Bill = EachBillToServe{}
				newVotesStruct.Bill.BillID = eachMembersBill.Bill.BillID
				newVotesStruct.Bill.Number = eachMembersBill.Bill.Number
				newVotesStruct.Bill.SponsorID = billsAPIResponse.Results[0].SponsorID
				newVotesStruct.Bill.BillURI = eachMembersBill.Bill.BillURI
				newVotesStruct.Bill.Title = eachMembersBill.Bill.Title
				newVotesStruct.Bill.LatestAction = eachMembersBill.Bill.LatestAction
				newVotesStruct.Bill.ShortTitle = billsAPIResponse.Results[0].ShortTitle
				newVotesStruct.Bill.PrimarySubject = billsAPIResponse.Results[0].PrimarySubject

				//Opensecrets Correlation
				newVotesStruct.Bill.OpensecretsSectorPrefix, newVotesStruct.Bill.OpensecretsSector, newVotesStruct.Bill.OpensecretsSectorLong, err = assign_opensecrets_data(billsAPIResponse.Results[0].PrimarySubject)
				if err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					http.Error(w, "Internal Server Error in Data Correlation", 400)
					return
				}

				//Bills API Response
				newVotesStruct.Bill.Committees = billsAPIResponse.Results[0].Committees
				newVotesStruct.Bill.CommitteeCodes = billsAPIResponse.Results[0].CommitteeCodes
				newVotesStruct.Bill.SubcommitteeCodes = billsAPIResponse.Results[0].SubcommitteeCodes
				newVotesStruct.Bill.Summary = billsAPIResponse.Results[0].Summary
				newVotesStruct.Bill.SummaryShort = billsAPIResponse.Results[0].SummaryShort

				//From Members API Again
				//newVotesStruct.Amendment = eachMembersBill.Amendment //REMOVED CAUSE WE DONT NEED IT
				newVotesStruct.Description = eachMembersBill.Description
				newVotesStruct.Question = eachMembersBill.Question
				newVotesStruct.Result = eachMembersBill.Result
				newVotesStruct.Date = eachMembersBill.Date
				newVotesStruct.Time = eachMembersBill.Time
				newVotesStruct.Total = eachMembersBill.Total
				newVotesStruct.Position = eachMembersBill.Position

				billsToServe.Votes = append(billsToServe.Votes, newVotesStruct)

				//fmt.Println("Primary subject: " + billsAPIResponse.Results[0].PrimarySubject)
			}
		}
	}

	//Finally, respond with a JSON-encoded version of the BillsToServe struct.
	enc := json.NewEncoder(w)
	if err := enc.Encode(billsToServe); err != nil {
		fmt.Printf("error encoding struct into JSON: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	//TODO - SOMETHIGN ABOUT CLOSING THE HTML STREAM TO AVOID WASTING RESOURCES

}

//TODO - MAKE THIS UNDERSTANDABLE AND NOT GODAWFULLY ORGANIZED
func isActuallyABill(question string, bill_id string) bool {
	switch question {
	case
		"On Motion to Suspend the Rules and Agree",
		"On Motion to Suspend the Rules and Pass, as Amended",
		"On Motion to Suspend the Rules and Pass",
		"On Passage",
		"On Agreeing to the Resolution",
		"On Passage of the Bill",
		"On the Joint Resolution",
		"On the Cloture Motion":

		switch bill_id {
		case
			"MOTION":
			return false
		}
		fmt.Println("Greenlit ID: " + bill_id)
		return true

	}
	fmt.Println("Rejected Question: " + question)
	return false
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

	log.Printf("Fetched bills recently voted on for" + member_id + "without errors, with length " + strconv.Itoa(len(responseData)))

	//Read the fetched data into the MemberAPIResponse struct
	membersAPIResponse := &MembersAPIResponse{}
	json.Unmarshal(responseData, membersAPIResponse)

	return membersAPIResponse, nil
}

//This function fetches bill details
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
		//fmt.Println("Primary subject: " + billsAPIResponse.Results[0].PrimarySubject)
		//fmt.Println("Short title: " + billsAPIResponse.Results[0].ShortTitle)
		//fmt.Println()
	} else {
		fmt.Println("No data returned for " + bill_slug)
	}

	return billsAPIResponse, nil
}

func assign_opensecrets_data(primary_subject string) (string, string, string, error) {
	switch primary_subject {
	case "Health":
		return "H", "Health", "Health", nil
	case "Government Operations and Politics":
		return "Z", "Joint Candidate Cmtes", "Joint Candidate Cmtes", nil //Anyone affected by this would be a politician
	case "International Affairs":
		return "Q", "Ideology/Single-Issue", "Ideological/Single-Issue", nil
	case "Congress":
		return "Z", "Joint Candidate Cmtes", "Joint Candidate Cmtes", nil //Anyone affected by this would be a politician
	case "Crime and Law Enforcement":
		return "P", "Labor", "Labor", nil //Police Unions fall under Labor in Opensecrets
	case "Taxation":
		return "Q", "Ideology/Single-Issue", "Ideological/Single-Issue", nil //Taxation falls under Single Issue in Opensecrets
	case "Armed Forces and National Security":
		return "D", "Defense", "Defense", nil
	case "Public Lands and Natural Resources":
		return "E", "Energy/Nat Resource", "Energy & Natural Resources", nil
	case "Education":
		return "W", "Other", "Other", nil
	case "Transportation and Public Works":
		return "M", "Transportation", "Transportation", nil
	case "Immigration":
		return "Q", "Ideology/Single-Issue", "Ideological/Single-Issue", nil
	case "Science, Technology, Communications":
		return "C", "Communic/Electronics", "Communications/Electronics", nil
	case "Labor and Employment":
		return "P", "Labor", "Labor", nil
	case "Commerce":
		return "N", "Misc Business", "Misc Business", nil
	case "Environmental Protection":
		return "Q", "Ideology/Single-Issue", "Ideological/Single-Issue", nil
	case "Finance and Financial Sector":
		return "F", "Finance/Insur/RealEst", "Finance, Insurance & Real Estate", nil
	case "Energy":
		return "E", "Energy/Nat Resource", "Energy & Natural Resources", nil
	case "Civil Rights and Liberties, Minority Issues":
		return "Q", "Ideology/Single-Issue", "Ideological/Single-Issue", nil
	case "Agriculture and Food":
		return "A", "Agribusiness", "Agribusiness", nil
	case "Native Americans":
		return "Q", "Ideology/Single-Issue", "Ideological/Single-Issue", nil
	case "Economics and Public Finance":
		return "F", "Finance/Insur/RealEst", "Finance, Insurance & Real Estate", nil
	case "Law":
		return "K", "Lawyers & Lobbyists", "Lawyers & Lobbyists", nil
	case "Housing and Community Development":
		return "C", "Construction", "Construction", nil
	case "Emergency Management":
		return "W", "Other", "Other", nil //Not sure what to put here, so...
	case "Social Welfare":
		return "P", "Labor", "Labor", nil //Most of the people affected by this are working-class
	case "Sports and Recreation":
		return "N", "Misc Business", "Misc Business", nil
	case "Foreign Trade and International Finance":
		return "N", "Misc Business", "Misc Business", nil
	case "Families":
		return "Y", "Unknown", "Unknown", nil //Homemakers are listed as Unknown
	case "Arts, Culture, Religion":
		return "W", "Other", "Other", nil
	case "Water Resources Development":
		return "E", "Energy/Nat Resource", "Energy & Natural Resources", nil
	case "Animals":
		return "A", "Agribusiness", "Agribusiness", nil

	}

	return "Z", "PROBLEM", "PROBLEM", nil //TODO add error handling to catch any that dont get sorted
}
