package handlers

import (
	"fmt"
	"net/http"
)

type BillsResponseStruct struct {
	Status  string `json:"status"`
	Results []struct {
		MemberID   string `json:"member_id"`
		NumResults string `json:"num_results"`
		Offset     string `json:"offset"`
		Votes      []struct {
			Session  string `json:"session"`
			RollCall string `json:"roll_call"`
			VoteURI  string `json:"vote_uri"`
			Bill     struct {
				BillID                  string        `json:"bill_id"`
				Number                  string        `json:"number"`
				SponsorID               string        `json:"sponsor_id"`
				BillURI                 string        `json:"bill_uri"`
				Title                   string        `json:"title"`
				LatestAction            string        `json:"latest_action"`
				ShortTitle              string        `json:"short_title"`
				PrimarySubject          string        `json:"primary_subject"`
				OpensecretsSectorPrefix string        `json:"Opensecrets_Sector_Prefix"`
				OpensecretsSector       string        `json:"Opensecrets_Sector"`
				OpensecretsSectorLong   string        `json:"Opensecrets_Sector_Long"`
				Committees              string        `json:"committees"`
				CommitteeCodes          []string      `json:"committee_codes"`
				SubcommitteeCodes       []interface{} `json:"subcommittee_codes"`
				Summary                 string        `json:"summary"`
				SummaryShort            string        `json:"summary_short"`
			} `json:"bill"`
			Amendment struct {
			} `json:"amendment"`
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
			MemberID string `json:"member_id,omitempty"`
		} `json:"votes"`
	} `json:"results"`
}

//DummyTopTenHandler returns Alma Adams's JSON data, and nobody else's. It does check that a name was provided
func DummyBillsHandler(w http.ResponseWriter, r *http.Request) {
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
			"status": "OK",
			"results": [{
				"member_id": "K000388",
				"num_results": "20",
				"offset": "0",
				"votes": [{
						"session": "1",
						"roll_call": "120",
						"vote_uri": "https://api.propublica.org/congress/v1/117/house/sessions/1/votes/120.json",
						"bill": {
							"bill_id": "hr1996-117",
							"number": "H R 1996",
							"sponsor_id": "P000593",
							"bill_uri": "https://api.propublica.org/congress/v1/117/bills/hr1996.json",
							"title": "To create protections for financial institutions that provide financial services to cannabis-related legitimate businesses and service providers for such businesses, and for other purposes.",
							"latest_action": "Referred to the Committee on Financial Services, and in addition to the Committee on the Judiciary, for a period to be subsequently determined by the Speaker, in each case for consideration of such provisions as fall within the jurisdiction of the committee concerned.",
							"short_title": "SAFE Banking Act of 2021",
							"primary_subject": "Finance and Financial Sector",
							"Opensecrets_Sector_Prefix": "F",
							"Opensecrets_Sector": "Finance/Insur/RealEst",
							"Opensecrets_Sector_Long": "Finance, Insurance & Real Estate",
							"committees": "House Judiciary Committee",
							"committee_codes": [
								"HSBA"
							],
							"subcommittee_codes": [],
							"summary": "Secure and Fair Enforcement Banking Act of 2021 or the SAFE Banking Act of 2021  This bill generally prohibits a federal banking regulator from penalizing a depository institution for providing banking services to a legitimate cannabis-related business. Prohibited penalties include terminating or limiting the deposit insurance or share insurance of a depository institution solely because the institution provides financial services to a legitimate cannabis-related business and prohibiting or otherwise discouraging a depository institution from offering financial services to such a business.   Additionally, proceeds from a transaction involving activities of a legitimate cannabis-related business are not considered proceeds from unlawful activity. Proceeds from unlawful activity are subject to anti-money laundering laws.  Furthermore, a depository institution is not, under federal law, liable or subject to asset forfeiture for providing a loan or other financial services to a legitimate cannabis-related business.  The bill also provides that a federal banking agency may not request or order a depository institution to terminate a customer account unless (1) the agency has a valid reason for doing so, and (2) that reason is not based solely on reputation risk. Valid reasons for terminating an account include threats to national security and involvement in terrorist financing, including state sponsorship of terrorism. ",
							"summary_short": "Secure and Fair Enforcement Banking Act of 2021 or the SAFE Banking Act of 2021  This bill generally prohibits a federal banking regulator from penalizing a depository institution for providing banking services to a legitimate cannabis-related business. Prohibited penalties include terminating or limiting the deposit insurance or share insurance of a depository institution solely because the institution provides financial services to a legitimate cannabis-related business and prohibiting or o..."
						},
						"amendment": {},
						"description": "SAFE Banking Act",
						"question": "On Motion to Suspend the Rules and Pass, as Amended",
						"result": "Passed",
						"date": "2021-04-19",
						"time": "19:05:00",
						"total": {
							"yes": 321,
							"no": 101,
							"present": 0,
							"not_voting": 8
						},
						"position": "No"
					},
					{
						"session": "1",
						"roll_call": "119",
						"vote_uri": "https://api.propublica.org/congress/v1/117/house/sessions/1/votes/119.json",
						"bill": {
							"bill_id": "hr1482-117",
							"number": "H R 1482",
							"sponsor_id": "B001311",
							"bill_uri": "https://api.propublica.org/congress/v1/117/bills/hr1482.json",
							"title": "To amend the Small Business Act to enhance the Office of Credit Risk Management, to require the Administrator of the Small Business Administration to issue rules relating to environmental obligations of certified development companies, and for other purposes.",
							"latest_action": "Received in the Senate and Read twice and referred to the Committee on Small Business and Entrepreneurship.",
							"short_title": "504 Credit Risk Management Improvement Act of 2021",
							"primary_subject": "Commerce",
							"Opensecrets_Sector_Prefix": "N",
							"Opensecrets_Sector": "Misc Business ",
							"Opensecrets_Sector_Long": "Misc Business "
						},
						"summary": "Secure and Fair Enforcement Banking Act of 2021 or the SAFE Banking Act of 2021  This bill generally prohibits a federal banking regulator from penalizing a depository institution for providing banking services to a legitimate cannabis-related business. Prohibited penalties include terminating or limiting the deposit insurance or share insurance of a depository institution solely because the institution provides financial services to a legitimate cannabis-related business and prohibiting or otherwise discouraging a depository institution from offering financial services to such a business.   Additionally, proceeds from a transaction involving activities of a legitimate cannabis-related business are not considered proceeds from unlawful activity. Proceeds from unlawful activity are subject to anti-money laundering laws.  Furthermore, a depository institution is not, under federal law, liable or subject to asset forfeiture for providing a loan or other financial services to a legitimate cannabis-related business.  The bill also provides that a federal banking agency may not request or order a depository institution to terminate a customer account unless (1) the agency has a valid reason for doing so, and (2) that reason is not based solely on reputation risk. Valid reasons for terminating an account include threats to national security and involvement in terrorist financing, including state sponsorship of terrorism. ",
						"summary_short": "Secure and Fair Enforcement Banking Act of 2021 or the SAFE Banking Act of 2021  This bill generally prohibits a federal banking regulator from penalizing a depository institution for providing banking services to a legitimate cannabis-related business. Prohibited penalties include terminating or limiting the deposit insurance or share insurance of a depository institution solely because the institution provides financial services to a legitimate cannabis-related business and prohibiting or o...",
	
						"amendment": {},
						"description": "504 Credit Risk Management Improvement Act",
						"question": "On Motion to Suspend the Rules and Pass",
						"result": "Passed",
						"date": "2021-04-16",
						"time": "12:19:00",
						"total": {
							"yes": 411,
							"no": 8,
							"present": 0,
							"not_voting": 11
						},
						"position": "Yes"
					},
					{
						"session": "1",
						"roll_call": "118",
						"vote_uri": "https://api.propublica.org/congress/v1/117/house/sessions/1/votes/118.json",
						"bill": {
							"bill_id": "hr1195-117",
							"number": "H R 1195",
							"sponsor_id": "C001069",
							"bill_uri": "https://api.propublica.org/congress/v1/117/bills/hr1195.json",
							"title": "To direct the Secretary of Labor to issue an occupational safety and health standard that requires covered employers within the health care and social service industries to develop and implement a comprehensive workplace violence prevention plan, and for other purposes.",
							"latest_action": "Received in the Senate and Read twice and referred to the Committee on Health, Education, Labor, and Pensions.",
							"short_title": "Workplace Violence Prevention for Health Care and Social Service Workers Act",
							"primary_subject": "Labor and Employment",
							"Opensecrets_Sector_Prefix": "P",
							"Opensecrets_Sector": "Labor",
							"Opensecrets_Sector_Long": "Labor",
							"committees": "House Judiciary Committee",
							"committee_codes": [
								"HSBA"
							],
							"subcommittee_codes": [],
							"summary": "Dummy long summary with fake data ",
							"summary_short": "Dummy short summary with fake data"
						},
						"amendment": {},
						"description": "Workplace Violence Prevention for Health Care and Social Service Workers Act",
						"question": "On Passage",
						"result": "Passed",
						"date": "2021-04-16",
						"time": "11:48:00",
						"total": {
							"yes": 254,
							"no": 166,
							"present": 0,
							"not_voting": 10
						},
						"position": "No"
					},
					{
						"member_id": "K000388",
						"session": "1",
						"roll_call": "116",
						"vote_uri": "https://api.propublica.org/congress/v1/117/house/sessions/1/votes/116.json",
						"bill": {
							"bill_id": "hr1490-117",
							"number": "H R 1490",
							"sponsor_id": "C001119",
							"bill_uri": "https://api.propublica.org/congress/v1/117/bills/hr1490.json",
							"title": "To amend the Small Business Investment Act of 1958 to improve the loan guaranty program, enhance the ability of small manufacturers to access affordable capital, and for other purposes.",
							"latest_action": "Received in the Senate and Read twice and referred to the Committee on Small Business and Entrepreneurship.",
							"short_title": " 504 Modernization and Small Manufacturer Enhancement Act of 2021",
							"primary_subject": "Commerce",
							"Opensecrets_Sector_Prefix": "N",
							"Opensecrets_Sector": "Misc Business ",
							"Opensecrets_Sector_Long": "Misc Business ",
							"committees": "House Judiciary Committee",
							"committee_codes": [
								"HSBA"
							],
							"subcommittee_codes": [],
							"summary": "Dummy long summary with fake data ",
							"summary_short": "Dummy short summary with fake data"
						},
						"amendment": {},
						"description": "504 Modernization and Small Manufacturer Enhancement Act",
						"question": "On Motion to Suspend the Rules and Pass",
						"result": "Passed",
						"date": "2021-04-15",
						"time": "21:42:00",
						"total": {
							"yes": 400,
							"no": 16,
							"present": 0,
							"not_voting": 14
						},
						"position": "Yes"
					},
					{
						"session": "1",
						"roll_call": "115",
						"vote_uri": "https://api.propublica.org/congress/v1/117/house/sessions/1/votes/115.json",
						"bill": {
							"bill_id": "hr1487-117",
							"number": "H R 1487",
							"sponsor_id": "B001309",
							"bill_uri": "https://api.propublica.org/congress/v1/117/bills/hr1487.json",
							"title": "To amend the Small Business Act to increase transparency, and for other purposes.",
							"latest_action": "Received in the Senate and Read twice and referred to the Committee on Small Business and Entrepreneurship.",
							"short_title": "Microloan Transparency and Accountability Act of 2021",
							"primary_subject": "Commerce",
							"Opensecrets_Sector_Prefix": "N",
							"Opensecrets_Sector": "Misc Business ",
							"Opensecrets_Sector_Long": "Misc Business ",
							"committees": "House Judiciary Committee",
							"committee_codes": [
								"HSBA"
							],
							"subcommittee_codes": [],
							"summary": "Dummy long summary with fake data ",
							"summary_short": "Dummy short summary with fake data"
						},
						"amendment": {},
						"description": "Microloan Transparency and Accountability Act",
						"question": "On Motion to Suspend the Rules and Pass",
						"result": "Passed",
						"date": "2021-04-15",
						"time": "21:10:00",
						"total": {
							"yes": 409,
							"no": 4,
							"present": 0,
							"not_voting": 17
						},
						"position": "Yes"
					}
				]
			}]
		}
	}`

	//  w.Write([]byte(jsonData))
	w.Write([]byte(jsonData))
}
