package useDB_for_handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

//dsn is a helper function. It returns the Data Source Name when db name passed as a parameter
func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

const (
	username = "root"
	password = "ftvInternal123" //FLAG this is a hardcoded password
	hostname = "mysqldb:3306"   //renaming this to be a docker container
	dbname   = "ftvBackEnd"
)

func Connect() (*sql.DB, error) {
	log.Printf("Entering the Connect funtion in package useDB_for_handler")

	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	//defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	//defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	return db, nil
}

//Creates the Top Ten table
func CreateTopTenTable(db *sql.DB) error {
	//query := `CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text,
	//    product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	query := `CREATE TABLE IF NOT EXISTS topten(id int not null auto_increment primary key,
		cand_name varchar(255) not null,
		cid varchar(64) not null unique,
		cycle varchar(64) not null,
		last_updated varchar(255) not null,
		last_updated_ftv_db datetime default CURRENT_TIMESTAMP,
	
		industry_code1 varchar(255) not null,
		industry_name1 varchar(255) not null,
		indivs1 int not null,
		pacs1 int not null,
		total1 int not null)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating topten table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}
