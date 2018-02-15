package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func listFiles(path string, db *sql.DB) error {
	var databasefileName string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		name := f.Name()
		// size := f.Size()
		//  fmt.Println(name)
		err := db.QueryRow("SELECT name FROM files WHERE name=? AND path=?", name, path).Scan(&databasefileName)
		// fmt.Println(&databasefileName)
		//There is no corresonding file in the directory
		if err != nil {
			_, err = db.Exec("INSERT INTO files(name, path) VALUES(?, ?)", name, path)
			if err != nil {
				fmt.Println("error in insertion: ")
				fmt.Println(err)
			}
		} else { //file as already been mapped in the directory
			// fmt.Println("already in db")
		}

		// fmt.Println(path + name)			// DEBUG

		file, err := os.Stat(path + name)
		if err != nil {
			return err
		}
		// in case file found is a directory it's calling listFiles recursively to index files in
		if file.IsDir() {
			fmt.Println("this is a directory : " + name)
			listFiles(path+name+"/", db)
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

/*FileIndexor goes to root of the computer and indexes every files in it
People just need to type any caracters in the search bar to find what they need
*/
func FileIndexor(db *sql.DB) error {
	err := listFiles("/", db)
	if err != nil {
		return err
	} else {
		fmt.Println("l'indexation s'est correctement déroulée")
		return nil
	}
}
