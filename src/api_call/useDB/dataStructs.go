package useDB

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
