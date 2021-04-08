package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
- read the csv and get a bunch of IDs
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

//Code
func main() {

	/*response, err := http.Get("http://www.opensecrets.org/api/?method=candIndustry&cid=N00007360&cycle=2020&apikey=c3fd74a75e5cb8756e262e8d2f0480b3")
	if err != nil {
		log.Printf("Error in fetching response")
		log.Fatal(err)
	} */

	crp_ids := readCSV()
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

			fmt.Println(congressperson.Response.Industries.Attributes.CandName)
			for _, industry := range congressperson.Response.Industries.Industry {
				fmt.Print(industry.Attributes.IndustryName + ", ")
			}
			fmt.Println("")
		}
	}
}

func readCSV() [][]string {

	file, err := os.Open("just_crp_ids.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	return records

}
