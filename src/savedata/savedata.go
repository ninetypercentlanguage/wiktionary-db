package savedata

import (
	"database/sql"
	"fmt"

	"github.com/ninetypercentlanguage/wiktionary-db/getdata"
)

func Save(w []getdata.WordData, db *sql.DB) {
	for i, word := range w {
		var wordId int
		row := db.QueryRow("INSERT INTO words (string) VALUES ($1) RETURNING id", word.Word)
		err := row.Scan(&wordId)
		if err != nil {
			panic(err)
		}

		for _, pos := range word.Content {
			partOfSpeechName := pos.PartOfSpeech
			var posId int
			row := db.QueryRow("INSERT INTO parts_of_speech (word, part_of_speech) VALUES ($1, $2) RETURNING id", wordId, partOfSpeechName)
			err := row.Scan(&posId)
			if err != nil {
				panic(err)
			}
			// todo lemmas table
		}
		if i%1000 == 0 {
			fmt.Printf("%v words\n", i)
		}
	}
}
