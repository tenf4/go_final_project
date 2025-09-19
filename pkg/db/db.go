package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `CREATE TABLE scheduler (
                    id INTEGER PRIMARY KEY AUTOINCREMENT, 
                    date CHAR(8) NOT NULL DEFAULT "",
                    title VARCHAR(255) NOT NULL,
                    comment TEXT,
                    repeat VARCHAR(128));`

var database *sql.DB

func Init(dbFile string) error {

	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	database, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	if install {
		_, err = database.Exec(schema)
		if err != nil {
			return err
		}
	}
	return nil
}
