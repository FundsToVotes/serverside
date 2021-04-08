package main

import (
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
*/
func main() {

	/*response, err := http.Get("http://www.opensecrets.org/api/?method=candIndustry&cid=N00007360&cycle=2020&apikey=c3fd74a75e5cb8756e262e8d2f0480b3")
	if err != nil {
		log.Printf("Error in fetching response")
		log.Fatal(err)
	} */

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://www.opensecrets.org/api/?method=candIndustry&cid=N00007360&cycle=2020&apikey=c3fd74a75e5cb8756e262e8d2f0480b3", nil)
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

	log.Printf("No error detected")
	fmt.Println(string(responseData))

}
