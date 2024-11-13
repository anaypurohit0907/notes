package data

import (
	"database/sql"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> f8a8337 (initial commit)
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS studybuddy (
        "idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "word" TEXT,
        "definition" TEXT,
        "category" TEXT
      );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Studybuddy table created")
}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO studybuddy(word, definition, category) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted study note successfully")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM studybuddy ORDER BY word")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string
		row.Scan(&idNote, &word, &definition, &category)
<<<<<<< HEAD
		fmt.Println("[", category, "] : ", word, "—", definition)
	}
}
func SearchAllNotes(word string) {
	row, err := db.Query("SELECT * FROM studybuddy WHERE definition LIKE '%" + word + "%'")

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string
		row.Scan(&idNote, &word, &definition, &category)
		fmt.Println("[", category, "] : ", word, "—", definition)
	}

}
=======
		log.Println("[", category, "] ", word, "—", definition)
	}
}
>>>>>>> f8a8337 (initial commit)
