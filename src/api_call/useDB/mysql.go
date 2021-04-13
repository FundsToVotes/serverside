package useDB

import "database/sql"

//MySQLStore is the struct for the sql store and is broken
type MySQLStore struct {
	conn *sql.DB
}

//NewMySQLStore creates a new mysqlstore object, but is broken
func NewMySQLStore(db *sql.DB) *MySQLStore {
	/*dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/demo", mysqlRootPW)
	 */
	return &MySQLStore{
		conn: db,
	}
}

//Inserts data into the DB
func Insert(cp *Congressperson) (*Congressperson, error) {
	//insert a new row into the "contacts" table
	//use ? markers for the values to defeat SQL
	//injection attacks
	insq := "insert into users(cand_name, cid, cycle, last_updated, last_updated_ftv_db, industry_code1, industry_name1, indivs1, pacs1, total1) values (?,?,?,?,?,?)"
	res, err := db.conn.Exec(insq, u.Email, u.PassHash, u.UserName, u.FirstName, u.LastName, u.PhotoURL)
	if err != nil {
		return nil, err
	}
	//get the auto-assigned ID for the new row
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	cp.ID = id
	return cp, nil
}
