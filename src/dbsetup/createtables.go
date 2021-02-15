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
			word            bigint NOT NULL,
			part_of_speech  varchar(30) NOT NULL,
			UNIQUE(word, part_of_speech),
			CONSTRAINT      fk_word
				FOREIGN KEY (word)
					REFERENCES words(id)
		);	
	`)
	if err != nil {
		return err
	}
	_, err = db.Query(`
		CREATE TABLE lemmas (
			id              bigserial PRIMARY KEY,
			part_of_speech	bigint NOT NULL,
			word			bigint NOT NULL,
			definitions		varchar(255),
			UNIQUE(word, part_of_speech),
			CONSTRAINT      fk_word
				FOREIGN KEY (word)
					REFERENCES words(id),
			CONSTRAINT      fk_part_of_speech
				FOREIGN KEY (part_of_speech)
					REFERENCES parts_of_speech(id)
		);	
	`)
	return err
}
