package dbhandler

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// Connect open connection to mysql instance
func Connect(dbName, userID, pwd string) error {
	database, err := sql.Open("mysql", userID+":"+pwd+"@/"+dbName)
	if err != nil {
		return err
	}
	db = database
	return nil
}

// Exec execute a query
func Exec(query string) (sql.Result, error) {
	db := db
	r, err := db.Exec(query)
	return r, err
}

// Query execute a query
func Query(query string) error {
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return nil
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return err
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)
	}
	return nil
}

// Db returns a pointer to mysql db
func Db() *sql.DB {
	if db == nil {
		return nil
	}
	return db
}

// Close close connection
func Close() {
	db.Close()
}
