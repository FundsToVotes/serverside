package main

import (
	"api_call/useDB"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Relevant tutorials
/*

https://tutorialedge.net/golang/consuming-restful-api-with-go/

https://www.soberkoder.com/consume-rest-api-go/

https://golangbot.com/connect-create-db-mysql/
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

/*Next Task

  + Inside APi folder
      > function that calls the relevant APIs and spits out data
          >> Inserts new things in the database
          >> Updates not-new things
      > Function for emptying database (developer use only)
*/
//Code
func main() {
	/*
		//Open a mysql database
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}
		//Initialize my user Store
		sqlStore := useDB.NewMySQLStore(db)
	*/

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

	//Managing the data

	openSecretsAPIKey := os.Getenv("SERVERSIDE_OPENSECRETS_API_KEY")

	if len(openSecretsAPIKey) == 0 {
		log.Fatal("Error - could not find SERVERSIDE_OPENSECRETS_API_KEY enviroment variable")
	}

	/*
		crp_ids := readCSV()

			for count, id := range crp_ids {
				//Skipped Count = 0 for now
				if 50 < count && count < 100 {
					log.Printf("Fetching and inserting crp_id " + id[0])

					client := &http.Client{}

					request_url := "http://www.opensecrets.org/api/?method=candIndustry&cid=" + id[0] + "&cycle=2020&apikey=" + openSecretsAPIKey + "&output=json"
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
						if strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
							log.Printf("Entry for member " + congressperson.Response.Industries.Attributes.CandName + " already exists, updating instead")
							//UPDATE FUNCTION TO DO
						} else {
							log.Printf("Insert member "+congressperson.Response.Industries.Attributes.CandName+" failed with error %s", err)
							return //TODO - DECIDE WHETHER OR NOT i WANT THE CODE TO RETURN UPON ONE PERSON ERRORING OUT
						}
					}

				}
			}
	*/

	//Retrieve data - This will not remain here in the long run, but is here for testing purposes
	/* migrated to handler code
	requested_CRP_Id := "N00031177" //choosing a random one of the crps loaded
	fmt.Println("CRP ID requested " + requested_CRP_Id)

	//Select from DB
	toptenToReturn, err := useDB.Select(db, requested_CRP_Id)
	if err != nil {
		log.Printf("Main: Error in Selecting "+requested_CRP_Id+", failed with error %s", err)
		return
	}

	fmt.Println("name to return: " + toptenToReturn.Attributes.CandName)
	fmt.Println("Industry1 to return: " + toptenToReturn.Industry[0].IndustryName)
	*/

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
