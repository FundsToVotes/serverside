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

	//Wow, this is ugly - I should've used an array
	response := &Top10Industriesresponse{}
	industry0 := &IndustryToServe{}
	industry1 := &IndustryToServe{}
	industry2 := &IndustryToServe{}
	industry3 := &IndustryToServe{}
	industry4 := &IndustryToServe{}
	industry5 := &IndustryToServe{}
	industry6 := &IndustryToServe{}
	industry7 := &IndustryToServe{}
	industry8 := &IndustryToServe{}
	industry9 := &IndustryToServe{}

	log.Println("Made it to the scan command")

	err = row.Scan(&response.Attributes.FtvDatabaseID, &response.Attributes.CandName, &response.Attributes.Cid, &response.Attributes.Cycle, &response.Attributes.LastUpdated,
		&response.Attributes.LastUpdatedFtVdb,
		&industry0.IndustryCode, &industry0.IndustryName, &industry0.Indivs, &industry0.Pacs, &industry0.Total,
		&industry1.IndustryCode, &industry1.IndustryName, &industry1.Indivs, &industry1.Pacs, &industry1.Total,
		&industry2.IndustryCode, &industry2.IndustryName, &industry2.Indivs, &industry2.Pacs, &industry2.Total,
		&industry3.IndustryCode, &industry3.IndustryName, &industry3.Indivs, &industry3.Pacs, &industry3.Total,
		&industry4.IndustryCode, &industry4.IndustryName, &industry4.Indivs, &industry4.Pacs, &industry4.Total,
		&industry5.IndustryCode, &industry5.IndustryName, &industry5.Indivs, &industry5.Pacs, &industry5.Total,
		&industry6.IndustryCode, &industry6.IndustryName, &industry6.Indivs, &industry6.Pacs, &industry6.Total,
		&industry7.IndustryCode, &industry7.IndustryName, &industry7.Indivs, &industry7.Pacs, &industry7.Total,
		&industry8.IndustryCode, &industry8.IndustryName, &industry8.Indivs, &industry8.Pacs, &industry8.Total,
		&industry9.IndustryCode, &industry9.IndustryName, &industry9.Indivs, &industry9.Pacs, &industry9.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			response = &Top10Industriesresponse{}
		} else {
			return nil, err
		}
	}

	response.Industry = append(response.Industry, *industry0)
	response.Industry = append(response.Industry, *industry1)
	response.Industry = append(response.Industry, *industry2)
	response.Industry = append(response.Industry, *industry3)
	response.Industry = append(response.Industry, *industry4)
	response.Industry = append(response.Industry, *industry5)
	response.Industry = append(response.Industry, *industry6)
	response.Industry = append(response.Industry, *industry7)
	response.Industry = append(response.Industry, *industry8)
	response.Industry = append(response.Industry, *industry9)

	return response, nil

}
