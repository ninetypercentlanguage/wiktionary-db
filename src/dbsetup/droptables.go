package dbsetup

import "database/sql"

func Drop(db *sql.DB) error {
	_, err := db.Query("DROP TABLE IF EXISTS lemmas;")
	if err != nil {
		return err
	}
	_, err = db.Query("DROP TABLE IF EXISTS parts_of_speech;")
	if err != nil {
		return err
	}
	_, err = db.Query("DROP TABLE IF EXISTS words;")
	return err
}
