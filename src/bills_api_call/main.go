package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	//_ "github.com/go-sql-driver/mysql"
)

// Relevant tutorials
/*
### Bills endpoint rambling and thoughts
Okay, so
I need to have several things
* API call to the propublica database that I want - which?? (chase this down soon)
    +  Call by bill, get category
* API call to the Opensecrets database
    + Call by person, get recent bills voted on?
* I think I'm slightly confused about what data comes from where - I want bill information, I want opensecrets information on the industry related to this, i want bills voted on by politician
* Goal: correlate donor industries (opensecrets) with bill categories (propublica), and use this to make a visualization
* Putting this together will require looking at a list of donor industries and a list of bill categories

1 - Get a Specific Member's Vote Positions
 - https://projects.propublica.org/api-docs/congress-api/members/.

2 - Getting the bill's category -
 this can actually also come from Propublica, using the primary_subject line of
 the Get a Specific Bill endpoint.
 https://projects.propublica.org/api-docs/congress-api/bills/
3  - Append the opensecrets industry category to the information about the bill,
 using this data - https://www.opensecrets.org/downloads/crp/CRP_Categories.txt
*/
//Code
func main() {
	log.Printf("The bills visualization!")
	member_id := "K000388"
	client := &http.Client{}

	request_url := "https://api.propublica.org/congress/v1/members/" + member_id + "/votes.json"
	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Golang_Funds_To_Votes_Bot")
	req.Header.Set("X-API-Key", "bUEZbt82MwpoNooSEbGjITvWKC703nY2isRQVa5Y")

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error in reading response")
		log.Fatal(err)
	}

	log.Printf("Data successfully fetched")

	//Create an empty instance of a MemberAPIResponse to recieve JSON
	membersAPIResponse := &MembersAPIResponse{}
	json.Unmarshal(responseData, membersAPIResponse)

	fmt.Println("~")
	fmt.Println(membersAPIResponse.Status)
	fmt.Println()

	for _, eachBill := range membersAPIResponse.Results[0].Votes {
		fmt.Println(eachBill.Bill.BillID + ": " /*+ eachBill.Bill.Title*/)
	}
	//fmt.Println("")

	fmt.Println("")
	fmt.Println("Now collecting specific data on each bill")

	debug_count := 1
	for index, eachBill := range membersAPIResponse.Results[0].Votes {
		if index <= debug_count {
			fmt.Println("")
			fmt.Println("fetching details for: " + eachBill.Bill.BillID)

			bill_slug := strings.TrimSuffix(eachBill.Bill.BillID, "-117")
			client := &http.Client{}

			request_url := "https://api.propublica.org/congress/v1/117/bills/" + bill_slug + ".json"
			req, err := http.NewRequest("GET", request_url, nil)
			if err != nil {
				log.Fatalln(err)
			}

			req.Header.Set("User-Agent", "Golang_Funds_To_Votes_Bot")
			req.Header.Set("X-API-Key", "bUEZbt82MwpoNooSEbGjITvWKC703nY2isRQVa5Y")

			response, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("Error in reading response")
				log.Fatal(err)
			}

			log.Printf("Bill data successfully fetched")

			//Create an empty instance of a BillsAPIResponse to recieve JSON
			billsAPIResponse := &BillsAPIResponse{}
			json.Unmarshal(responseData, billsAPIResponse)

			fmt.Println("~")
			fmt.Println(billsAPIResponse.Results[0].PrimarySubject)
			fmt.Println()
		}
	}
}

type MembersAPIResponse struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
	Results   []struct {
		MemberID   string `json:"member_id"`
		TotalVotes string `json:"total_votes"`
		Offset     string `json:"offset"`
		Votes      []struct {
			MemberID string `json:"member_id"`
			Chamber  string `json:"chamber"`
			Congress string `json:"congress"`
			Session  string `json:"session"`
			RollCall string `json:"roll_call"`
			VoteURI  string `json:"vote_uri"`
			Bill     struct {
				BillID       string `json:"bill_id"`
				Number       string `json:"number"`
				BillURI      string `json:"bill_uri"`
				Title        string `json:"title"`
				LatestAction string `json:"latest_action"`
			} `json:"bill,omitempty"`
			Description string `json:"description"`
			Question    string `json:"question"`
			Result      string `json:"result"`
			Date        string `json:"date"`
			Time        string `json:"time"`
			Total       struct {
				Yes       int `json:"yes"`
				No        int `json:"no"`
				Present   int `json:"present"`
				NotVoting int `json:"not_voting"`
			} `json:"total"`
			Position string `json:"position"`
			//Bill     struct {
			//} `json:"bill,omitempty"`
			//Bill struct {
			//} `json:"bill,omitempty"`
		} `json:"votes"`
	} `json:"results"`
}

type BillsAPIResponse struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
	Results   []struct {
		BillID            string      `json:"bill_id"`
		BillSlug          string      `json:"bill_slug"`
		Congress          string      `json:"congress"`
		Bill              string      `json:"bill"`
		BillType          string      `json:"bill_type"`
		Number            string      `json:"number"`
		BillURI           string      `json:"bill_uri"`
		Title             string      `json:"title"`
		ShortTitle        string      `json:"short_title"`
		SponsorTitle      string      `json:"sponsor_title"`
		Sponsor           string      `json:"sponsor"`
		SponsorID         string      `json:"sponsor_id"`
		SponsorURI        string      `json:"sponsor_uri"`
		SponsorParty      string      `json:"sponsor_party"`
		SponsorState      string      `json:"sponsor_state"`
		GpoPdfURI         interface{} `json:"gpo_pdf_uri"`
		CongressdotgovURL string      `json:"congressdotgov_url"`
		GovtrackURL       string      `json:"govtrack_url"`
		IntroducedDate    string      `json:"introduced_date"`
		Active            bool        `json:"active"`
		LastVote          interface{} `json:"last_vote"`
		HousePassage      string      `json:"house_passage"`
		SenatePassage     interface{} `json:"senate_passage"`
		Enacted           interface{} `json:"enacted"`
		Vetoed            interface{} `json:"vetoed"`
		Cosponsors        int         `json:"cosponsors"`
		CosponsorsByParty struct {
			R int `json:"R"`
			D int `json:"D"`
		} `json:"cosponsors_by_party"`
		WithdrawnCosponsors   int           `json:"withdrawn_cosponsors"`
		PrimarySubject        string        `json:"primary_subject"`
		Committees            string        `json:"committees"`
		CommitteeCodes        []string      `json:"committee_codes"`
		SubcommitteeCodes     []interface{} `json:"subcommittee_codes"`
		LatestMajorActionDate string        `json:"latest_major_action_date"`
		LatestMajorAction     string        `json:"latest_major_action"`
		HousePassageVote      string        `json:"house_passage_vote"`
		SenatePassageVote     interface{}   `json:"senate_passage_vote"`
		Summary               string        `json:"summary"`
		SummaryShort          string        `json:"summary_short"`
		Versions              []struct {
			Status            string `json:"status"`
			Title             string `json:"title"`
			URL               string `json:"url"`
			CongressdotgovURL string `json:"congressdotgov_url"`
		} `json:"versions"`
		Actions []struct {
			ID          int    `json:"id"`
			Chamber     string `json:"chamber"`
			ActionType  string `json:"action_type"`
			Datetime    string `json:"datetime"`
			Description string `json:"description"`
		} `json:"actions"`
		Votes []interface{} `json:"votes"`
	} `json:"results"`
}

//ReadCSV reads csvs
func readCSV() [][]string {

	file, err := os.Open("just_crp_ids.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	return records

}
