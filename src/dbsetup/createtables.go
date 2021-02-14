package dbsetup

import "database/sql"

func Create(db *sql.DB) error {
	_, err := db.Query(`
		CREATE TABLE words (
			id              bigserial PRIMARY KEY,
			string          varchar(50) NOT NULL,
			UNIQUE(string)
		);	
	`)
	if err != nil {
		return err
	}
	_, err = db.Query(`
		CREATE TABLE parts_of_speech (
			id              bigserial PRIMARY KEY,
			word            bigint,
			part_of_speech  varchar(30) NOT NULL,
			UNIQUE(word, part_of_speech),
			CONSTRAINT      fk_word
				FOREIGN KEY (word)
					REFERENCES words(id)
		);	
	`)
	return err
}
