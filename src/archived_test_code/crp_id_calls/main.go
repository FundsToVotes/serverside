package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//_ "github.com/go-sql-driver/mysql"
)

/*
1 - write some code that calls my api and creates the json file
TODO - ADD A SENATE CALL
2 - write some code that reads the json file into a map
3 - stick that in my handlers
*/
//Code
func main() {
	log.Printf("The bills visualization!")

	client := &http.Client{}

	request_url := "https://api.propublica.org/congress/v1/117/house/members.json"
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
	allMembersResponse := &AllMembersResponse{}
	json.Unmarshal(responseData, allMembersResponse)

	crpIDsToEncode := &CrpIDsToEncode{}

	for _, eachMember := range allMembersResponse.Results[0].Members {
		fmt.Println(eachMember.LastName+", "+eachMember.FirstName+" "+eachMember.MiddleName+": ", eachMember.CrpID)

		newPerson := Person{}

		newPerson.FirstName = eachMember.FirstName
		newPerson.MiddleName = eachMember.MiddleName
		newPerson.LastName = eachMember.LastName
		newPerson.CrpID = eachMember.CrpID

		if eachMember.CrpID == "" {
			fmt.Println(eachMember)

		}

		crpIDsToEncode.EachID = append(crpIDsToEncode.EachID, newPerson)
	}
	jsonToFile, _ := json.MarshalIndent(crpIDsToEncode, "", " ")

	_ = ioutil.WriteFile("crp_ids.json", jsonToFile, 0644)

	/*
		file, _ := os.OpenFile("crp_ids.json", os.O_CREATE, os.ModePerm)
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode(crpIDsToEncode)   */
}

//What we should encode
type CrpIDsToEncode struct {
	EachID []Person `json:"Person"`
}

//A struct for each new person
type Person struct {
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"Lastname"`
	CrpID      string `json:"CrpID"`
}

//Response from the Members API
type AllMembersResponse struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
	Results   []struct {
		Congress   string `json:"congress"`
		Chamber    string `json:"chamber"`
		NumResults int    `json:"num_results"`
		Offset     int    `json:"offset"`
		Members    []struct {
			ID                   string      `json:"id"`
			Title                string      `json:"title"`
			ShortTitle           string      `json:"short_title"`
			APIURI               string      `json:"api_uri"`
			FirstName            string      `json:"first_name"`
			MiddleName           string      `json:"middle_name"` //was an interface for some reason
			LastName             string      `json:"last_name"`
			Suffix               interface{} `json:"suffix"`
			DateOfBirth          string      `json:"date_of_birth"`
			Gender               string      `json:"gender"`
			Party                string      `json:"party"`
			LeadershipRole       interface{} `json:"leadership_role"`
			TwitterAccount       string      `json:"twitter_account"`
			FacebookAccount      string      `json:"facebook_account"`
			YoutubeAccount       string      `json:"youtube_account"`
			GovtrackID           string      `json:"govtrack_id"`
			CspanID              string      `json:"cspan_id"`
			VotesmartID          string      `json:"votesmart_id"`
			IcpsrID              string      `json:"icpsr_id"`
			CrpID                string      `json:"crp_id"`
			GoogleEntityID       string      `json:"google_entity_id"`
			FecCandidateID       string      `json:"fec_candidate_id"`
			URL                  string      `json:"url"`
			RssURL               string      `json:"rss_url"`
			ContactForm          string      `json:"contact_form"`
			InOffice             bool        `json:"in_office"`
			CookPvi              interface{} `json:"cook_pvi"`
			DwNominate           float64     `json:"dw_nominate"`
			IdealPoint           interface{} `json:"ideal_point"`
			Seniority            string      `json:"seniority"`
			NextElection         string      `json:"next_election"`
			TotalVotes           int         `json:"total_votes"`
			MissedVotes          int         `json:"missed_votes"`
			TotalPresent         int         `json:"total_present"`
			LastUpdated          string      `json:"last_updated"`
			OcdID                string      `json:"ocd_id"`
			Office               string      `json:"office"`
			Phone                string      `json:"phone"`
			Fax                  string      `json:"fax"`
			State                string      `json:"state"`
			SenateClass          string      `json:"senate_class"`
			StateRank            string      `json:"state_rank"`
			LisID                string      `json:"lis_id"`
			MissedVotesPct       float64     `json:"missed_votes_pct"`
			VotesWithPartyPct    float64     `json:"votes_with_party_pct"`
			VotesAgainstPartyPct float64     `json:"votes_against_party_pct"`
		} `json:"members"`
	} `json:"results"`
}
