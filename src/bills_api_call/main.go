package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		fmt.Println(eachBill.Bill.BillID + ": " + eachBill.Bill.Title)
	}
	//fmt.Println("")

	/*
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
	*/

	/*
		//Managing the data
		crp_ids := readCSV()
		for count, id := range crp_ids {
			//Skipped Count = 0 for now
			if 7 < count && count < 50 {
				log.Printf("Fetching and inserting crp_id " + id[0])

				client := &http.Client{}

				request_url := "http://www.opensecrets.org/api/?method=candIndustry&cid=" + id[0] + "&cycle=2020&apikey=c3fd74a75e5cb8756e262e8d2f0480b3&output=json"
				req, err := http.NewRequest("GET", request_url, nil)
				if err != nil {
					log.Fatalln(err)
				}

				req.Header.Set("User-Agent", "Golang_Funds_To_Votes_Bot")

				response, err := client.Do(req)
				if err != nil {
					log.Fatalln(err)
				}

				responseData, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Printf("Error in reading response")
					log.Fatal(err)
				}

				//Create an empty instance of a Congressperson to recieve JSON
				congressperson := &useDB.Congressperson{}
				json.Unmarshal(responseData, congressperson)

				fmt.Println("~")
				fmt.Println(congressperson.Response.Industries.Attributes.CandName)
				fmt.Println()

				fmt.Println("Raw Json")
				// Convert response body to string
				bodyString := string(responseData)
				fmt.Println(bodyString)
				fmt.Println("")

				for _, industry := range congressperson.Response.Industries.Industry {
					fmt.Print(industry.Attributes.IndustryName + ", ")
				}
				fmt.Println("")

				//Inserting into the db
				err = useDB.Insert(db, *congressperson)
				if err != nil {
					log.Printf("Insert member"+congressperson.Response.Industries.Attributes.CandName+" failed with error %s", err)
					return
				}

			}
		}

	*/
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
