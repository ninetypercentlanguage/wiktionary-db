package savedata

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ninetypercentlanguage/wiktionary-db/getdata"
	"github.com/ninetypercentlanguage/word-utils/combined"
)

type lemmaRefs struct {
	partOfSpeechId int
	lemmaData      combined.LemmaItem
}

func Save(w []getdata.WordData, db *sql.DB) {
	var idsForLemmas []lemmaRefs

	for i, word := range w {
		wordId := saveWord(word.Word, db)
		for _, pos := range word.Content {
			posId := savePartOfSpeech(pos, wordId, db)

			for _, lemma := range pos.Lemmas {
				idsForLemmas = append(idsForLemmas, lemmaRefs{
					partOfSpeechId: posId,
					lemmaData:      lemma,
				})
			}
		}
		if i%1000 == 0 {
			fmt.Printf("%v of %v words (and parts of speech) saved\n", i, len(w))
		}
	}

	for _, lemma := range idsForLemmas {
		saveLemma(lemma.lemmaData, lemma.partOfSpeechId, db)
	}
}

func saveWord(word string, db *sql.DB) int {
	var wordId int
	row := db.QueryRow("INSERT INTO words (string) VALUES ($1) RETURNING id", word)
	err := row.Scan(&wordId)
	if err != nil {
		panic(err)
	}
	return wordId
}

func savePartOfSpeech(pos combined.ContentItem, wordId int, db *sql.DB) int {
	partOfSpeechName := pos.PartOfSpeech
	var posId int
	row := db.QueryRow("INSERT INTO parts_of_speech (word, part_of_speech) VALUES ($1, $2) RETURNING id", wordId, partOfSpeechName)
	err := row.Scan(&posId)
	if err != nil {
		panic(err)
	}
	return posId
}

func saveLemma(lemma combined.LemmaItem, partOfSpeechId int, db *sql.DB) {
	var wordId int
	row := db.QueryRow("SELECT id FROM words WHERE string=$1", lemma.Word)
	err := row.Scan(&wordId)
	if err != nil {
		fmt.Println("partofSpeechid", partOfSpeechId)
		fmt.Println("lemma")
		panic(err)
	}
	definitions := lemma.Definitions
	if len(definitions) > 10 {
		definitions = definitions[:10]
	}
	rows, err := db.Query(
		"INSERT INTO lemmas (part_of_speech, word, definitions) VALUES ($1, $2, $3)",
		partOfSpeechId,
		wordId,
		strings.Join(definitions, ","),
	)
	if err != nil {
		panic(err)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}
}
