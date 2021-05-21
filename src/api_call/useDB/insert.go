package useDB

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

func Insert(db *sql.DB, cp Congressperson) error {
	//Remove party affiliation from name
	trimmedName := strings.Trim(strings.Split(cp.Response.Industries.Attributes.CandName, "(")[0], " ")

	query := "INSERT INTO topten(cand_name, cid, cycle, last_updated, last_updated_ftv_db, industry_code1, industry_name1, indivs1, pacs1, total1) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
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

		res, err := stmt.ExecContext(ctx, trimmedName, cpA.Cid, cpA.Cycle, cpA.LastUpdated, time.Now(), cpI[0].Attributes.IndustryCode, cpI[0].Attributes.IndustryName, cpI[0].Attributes.Indivs, cpI[0].Attributes.Pacs, cpI[0].Attributes.Total)
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
	}
	return nil
}
