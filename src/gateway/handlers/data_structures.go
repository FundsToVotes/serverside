package handlers

//The data that this custom backend will return for the Bills endpoint
type BillsToServe struct {
	Status string `json:"status"`
	//Results []struct {
	MemberID   string         `json:"member_id"`
	NumResults string         `json:"num_results"`
	TotalVotes string         `json:"total_votes"`
	Offset     string         `json:"offset"`
	Votes      []VotesToServe `json:"votes"`
	//} `json:"results"`
}

//Votes struct to be included inside BillsToServe
type VotesToServe struct {
	Session  string `json:"session"`
	RollCall string `json:"roll_call"`
	VoteURI  string `json:"vote_uri"`

	Bill EachBillToServe `json:"bill"`
	//Amendment struct {
	//} `json:"amendment"`

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
}

// The data for each individual bill being served
type EachBillToServe struct {
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
}

//The data fetched from the Propublica members endpoint
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

//The data fetched from the Bills propublica API endpoint
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
