package main

import (
	"database/sql"
	"fmt"

	_ "github.com/alexbrainman/odbc"
)

type Account struct {
	Id            string
	Name          string
	AccountNumber string
}

func main() {
	db, err := sql.Open("odbc", "DSN=CData Salesforce Source")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT Id, Name, AccountNumber FROM Account Limit 10")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var account Account

		rows.Scan(&account.Id, &account.Name, &account.AccountNumber)
		fmt.Printf("%+v\n", account)
	}
}
