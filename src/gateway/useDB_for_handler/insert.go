package useDB_for_handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Insert(db *sql.DB, cp Congressperson) error {
	//Remove party affiliation from name
	trimmedName := strings.Trim(strings.Split(cp.Response.Industries.Attributes.CandName, "(")[0], " ")

	//Making each congressperson have 10 industries

	// Creating and initializing anonymous struct
	/*emptyIndustry := struct {
		IndustryCode string
		IndustryName string
		Indivs       string
		Pacs         string
		Total        string
	}{
		IndustryCode: "",
		IndustryName: "",
		Indivs:       "",
		Pacs:         "",
		Total:        "",
	} */

	log.Println(strconv.Itoa(len(cp.Response.Industries.Industry)) + " industries total for this person")
	log.Println("")

	fmt.Printf("\n%+v\n", cp)
	log.Printf("")

	query := "INSERT INTO topten(cand_name, cid, cycle, last_updated, last_updated_ftv_db, " +
		"industry_code0, industry_name0, indivs0, pacs0, total0, " +
		"industry_code1, industry_name1, indivs1, pacs1, total1, industry_code2, industry_name2, indivs2, pacs2, total2, " +
		"industry_code3, industry_name3, indivs3, pacs3, total3, industry_code4, industry_name4, indivs4, pacs4, total4, " +
		"industry_code5, industry_name5, indivs5, pacs5, total5, industry_code6, industry_name6, indivs6, pacs6, total6, " +
		"industry_code7, industry_name7, indivs7, pacs7, total7, industry_code8, industry_name8, indivs8, pacs8, total8, " +
		"industry_code9, industry_name9, indivs9, pacs9, total9)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	log.Println(query)

	log.Println("")

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	cpA := cp.Response.Industries.Attributes
	cpI := cp.Response.Industries.Industry

	log.Printf("Adding the following industries for candidate" + cpA.CandName)
	if len(cpI) > 0 { //TODO MAKE THIS AN EMPTY STRUCT INSTEAD
		for _, industry := range cpI {
			fmt.Print(industry.Attributes.IndustryName + ", ")
		}
		fmt.Println("")
		fmt.Println("Got to the exec context line")

		res, err := stmt.ExecContext(ctx, trimmedName, cpA.Cid, cpA.Cycle, cpA.LastUpdated, time.Now(),
			cpI[0].Attributes.IndustryCode, cpI[0].Attributes.IndustryName, cpI[0].Attributes.Indivs, cpI[0].Attributes.Pacs, cpI[0].Attributes.Total,
			cpI[1].Attributes.IndustryCode, cpI[1].Attributes.IndustryName, cpI[1].Attributes.Indivs, cpI[1].Attributes.Pacs, cpI[1].Attributes.Total,
			cpI[2].Attributes.IndustryCode, cpI[2].Attributes.IndustryName, cpI[2].Attributes.Indivs, cpI[2].Attributes.Pacs, cpI[2].Attributes.Total,
			cpI[3].Attributes.IndustryCode, cpI[3].Attributes.IndustryName, cpI[3].Attributes.Indivs, cpI[3].Attributes.Pacs, cpI[3].Attributes.Total,
			cpI[4].Attributes.IndustryCode, cpI[4].Attributes.IndustryName, cpI[4].Attributes.Indivs, cpI[4].Attributes.Pacs, cpI[4].Attributes.Total,
			cpI[5].Attributes.IndustryCode, cpI[5].Attributes.IndustryName, cpI[5].Attributes.Indivs, cpI[5].Attributes.Pacs, cpI[5].Attributes.Total,
			cpI[6].Attributes.IndustryCode, cpI[6].Attributes.IndustryName, cpI[6].Attributes.Indivs, cpI[6].Attributes.Pacs, cpI[6].Attributes.Total,
			cpI[7].Attributes.IndustryCode, cpI[7].Attributes.IndustryName, cpI[7].Attributes.Indivs, cpI[7].Attributes.Pacs, cpI[7].Attributes.Total,
			cpI[8].Attributes.IndustryCode, cpI[8].Attributes.IndustryName, cpI[8].Attributes.Indivs, cpI[8].Attributes.Pacs, cpI[8].Attributes.Total,
			cpI[9].Attributes.IndustryCode, cpI[9].Attributes.IndustryName, cpI[9].Attributes.Indivs, cpI[9].Attributes.Pacs, cpI[9].Attributes.Total)
		if err != nil {
			log.Printf("Error %s when inserting row into topTen table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d congressperson records created ", rows)
	} else {
		res, err := stmt.ExecContext(ctx, trimmedName, cpA.Cid, cpA.Cycle, cpA.LastUpdated, time.Now(),
			"", "", "", "", "",
			"", "", "", "", "",
			"", "", "", "", "",
			"", "", "", "", "",
			"", "", "", "", "",
			"", "", "", "", "",
			"", "", "", "", "",
			"", "", "", "", "",
			"", "", "", "", "",
			"", "", "", "", "")
		if err != nil {
			log.Printf("Error %s when inserting row into topTen table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d NULL congressperson records created ", rows)

	}
	return nil
}
