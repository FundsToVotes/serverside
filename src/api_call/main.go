package main

import (
	"api_call/fetchdata"
	"api_call/useDB"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Relevant tutorials
/*

https://tutorialedge.net/golang/consuming-restful-api-with-go/

https://www.soberkoder.com/consume-rest-api-go/
*/

/* Dev notes
- apparently i needed a user-agent, and got a cloudflare error when I didnt have it
CRP IDs: - N00007360 for Nancy Pelosi

Process:
- x____read the csv and get a bunch of IDs
- x____make the api request
- x___store the api request in some structs
- store the struct content in a DB
*/

//Data structures
type Congressperson struct {
	Response struct {
		Industries struct {
			Attributes struct {
				CandName    string `json:"cand_name"`
				Cid         string `json:"cid"`
				Cycle       string `json:"cycle"`
				Origin      string `json:"origin"`
				Source      string `json:"source"`
				LastUpdated string `json:"last_updated"`
			} `json:"@attributes"`
			Industry []struct {
				Attributes struct {
					IndustryCode string `json:"industry_code"`
					IndustryName string `json:"industry_name"`
					Indivs       string `json:"indivs"`
					Pacs         string `json:"pacs"`
					Total        string `json:"total"`
				} `json:"@attributes"`
			} `json:"industry"`
		} `json:"industries"`
	} `json:"response"`
}

/*Next Task

  + Inside APi folder
      > function that calls the relevant APIs and spits out data
          >> Inserts new things in the database
          >> Updates not-new things
      > Function for emptying database (developer use only)
*/
//Code
func main() {
	//Open a mysql database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	//Initialize my user Store
	sqlStore := useDB.NewMySQLStore(db)

	crp_ids := fetchdata.ReadCSV()
	for count, id := range crp_ids {
		if count < 5 {
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
			congressperson := &Congressperson{}
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
			sqlStore.Insert(congressperson)
		}
	}
}
