package handlers

import (
	"fmt"
	"net/http"
)

type ResponseStruct struct {
	Top10Industriesresponse struct {
		Attributes struct {
			CandName                  string `json:"cand_name"`
			Cid                       string `json:"cid"`
			Cycle                     string `json:"cycle"`
			LastUpdated               string `json:"last_updated"`
			LastUpdatedFundstovotesdb string `json:"last_updated_FundsToVotesdb"`
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
	} `json:"top10IndustriesResponse"`
}

//DummyTopTenHandler returns Alma Adams's JSON data, and nobody else's. It does check that a name was provided
func DummyTopTenHandler(w http.ResponseWriter, r *http.Request) {
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

	jsonData := `{
			"top10IndustriesResponse": {
			  "@attributes": {
				"cand_name": "Alma Adams",
				"cid": "N00035451",
				"cycle": "2020",
				"last_updated": "03/22/2021",
				"last_updated_FundsToVotesdb": "04/09/2021"
			  },
			  "industry": [
				{
				  "@attributes": {
					"industry_code": "F03",
					"industry_name": "Commercial Banks",
					"indivs": "7178",
					"pacs": "58000",
					"total": "65178"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "F09",
					"industry_name": "Insurance",
					"indivs": "251",
					"pacs": "57000",
					"total": "57251"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "P04",
					"industry_name": "Public Sector Unions",
					"indivs": "750",
					"pacs": "51000",
					"total": "51750"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "F10",
					"industry_name": "Real Estate",
					"indivs": "17700",
					"pacs": "23500",
					"total": "41200"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "F07",
					"industry_name": "Securities & Investment",
					"indivs": "12",
					"pacs": "37000",
					"total": "37012"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "P02",
					"industry_name": "Industrial Unions",
					"indivs": "0",
					"pacs": "35500",
					"total": "35500"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "H01",
					"industry_name": "Health Professionals",
					"indivs": "19097",
					"pacs": "15500",
					"total": "34597"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "F06",
					"industry_name": "Finance/Credit Companies",
					"indivs": "0",
					"pacs": "27500",
					"total": "27500"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "K01",
					"industry_name": "Lawyers/Law Firms",
					"indivs": "10330",
					"pacs": "15000",
					"total": "25330"
				  }
				},
				{
				  "@attributes": {
					"industry_code": "A01",
					"industry_name": "Crop Production & Basic Processing",
					"indivs": "0",
					"pacs": "25000",
					"total": "25000"
				  }
				}
			  ]
			}
		  }`

	//  w.Write([]byte(jsonData))
	w.Write([]byte(jsonData))
}
