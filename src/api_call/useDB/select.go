package useDB

import (
	"database/sql"
	"log"
)

type Top10Industriesresponse struct {
	Attributes AttributesToServe `json:"@attributes"`
	Industry   []IndustryToServe `json:"industry"`
}

type AttributesToServe struct {
	FtvDatabaseID    string `json:"ftv_database_id"`
	CandName         string `json:"cand_name"`
	Cid              string `json:"cid"`
	Cycle            string `json:"cycle"`
	LastUpdated      string `json:"last_updated"`
	LastUpdatedFtVdb string `json:"last_updated_FundsToVotesdb"`
}

type IndustryToServe struct {
	IndustryCode string `json:"industry_code"`
	IndustryName string `json:"industry_name"`
	Indivs       string `json:"indivs"`
	Pacs         string `json:"pacs"`
	Total        string `json:"total"`
}

func Select(db *sql.DB, crp_id string) (*Top10Industriesresponse, error) {
	err := db.Ping()
	if err != nil {
		return &Top10Industriesresponse{}, err
	}

	query := "SELECT * FROM topten WHERE cid = ?"

	row := db.QueryRow(query, crp_id)

	response := &Top10Industriesresponse{}
	industry1 := &IndustryToServe{}

	log.Println("Made it to the scan command")

	err = row.Scan(&response.Attributes.FtvDatabaseID, &response.Attributes.CandName, &response.Attributes.Cid, &response.Attributes.Cycle, &response.Attributes.LastUpdated,
		&response.Attributes.LastUpdatedFtVdb, &industry1.IndustryCode, &industry1.IndustryName, &industry1.Indivs, &industry1.Pacs, &industry1.Total) //Doesn't have the actual industries
	if err != nil {
		if err == sql.ErrNoRows {
			response = &Top10Industriesresponse{}
		} else {
			return nil, err
		}
	}

	response.Industry = append(response.Industry, *industry1)

	return response, nil

}
