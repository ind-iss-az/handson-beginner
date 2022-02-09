package main

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/google/uuid"
)

var Db *sql.DB

type DbSetting struct {
	dsn string
}

func newDbSetting(param *Parameter) *DbSetting {
	ds := new(DbSetting)
	db_name := "azurestudy"
	ds.dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",param.db_host, param.db_user, param.db_password, 1433, db_name)
	return ds
}

type Message struct {
    MessageID   string	`db:"MessageID"`
    Date 		string	`db:"Date"`
    Hostname	string	`db:"Hostname"`
}

func PutMessages(Date string, Hostname string) (messages []Message) {

	uuidObj, _ := uuid.NewRandom()
	q_insert := fmt.Sprintf("INSERT INTO Message (MessageID, Date, Hostname) VALUES ('%s', '%s', '%s')", uuidObj.String(), Date, Hostname)
	q_select := fmt.Sprintf("SELECT * FROM Message WHERE MessageID = '%s'",uuidObj.String())

	tx, err := Db.Begin()
	tx.Exec(q_insert)
	tx.Commit()

	var readMessages []Message
	//var rows *sql.Row
	rows, err := Db.Query(q_select)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var r Message
		err = rows.Scan(&r.MessageID, &r.Date, &r.Hostname)
		if err != nil {
			log.Fatal(err)
		}
		readMessages = append(readMessages, r)
	}
	return readMessages
}