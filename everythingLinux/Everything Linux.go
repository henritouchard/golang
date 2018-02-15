package main

import (
	// "fileindexor"
	"net/http"
	//"html/template"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/henri/fileindexor"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func queryFile(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Println("it's not a get request")
		http.ServeFile(res, req, "searchFile.html")
		return
	}

	fileName := req.FormValue("FileName")
	fmt.Println("FileName = " + fileName)
	// anyChar := req.FormValue("anyChar")

	// var dbFileName string
	// var dbPath string

	rows, err := db.Query("SELECT name, path FROM files WHERE name LIKE ?", "%"+fileName+"%") //.Scan(&dbFileName, &dbPath)
	defer rows.Close()
	var resultName string
	var filePath string
	var result string

	for rows.Next() {
		err := rows.Scan(&resultName, &filePath)
		if err != nil {
			fmt.Println(err)
		}
		result += resultName + " in path : =====> " + filePath + "\n"

	}
	if err != nil {
		fmt.Println("query error")
		fmt.Println(err)
		http.Redirect(res, req, "/", 301)
		return
	}
	fmt.Println("ok query done")
	// var result string
	// for fi range := rows {
	// 	result += fi
	// 	result += '\n'
	// }

	res.Write([]byte(result))
}

// func createTable() {
// 	_, err := db.Exec("CREATE TABLE IF NOT EXISTS files(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name varchar(500), path varchar(1000)))")
// 	if err != nil{
// 		panic(err.Error())
// 	}
// }

func main() {

	// Create an sql.DB and check for errors
	// createTable()
	db, err = sql.Open("mysql", "root:T]x;4bghC2@/everything")
	if err != nil {
		panic(err.Error())
	}

	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()
	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	go fileindexor.FileIndexor(db)
	http.HandleFunc("/", queryFile)
	http.ListenAndServe(":3040", nil)
}
