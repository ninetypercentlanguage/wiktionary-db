package main

import (
	"database/sql"
	"log"

	"github.com/ninetypercentlanguage/wiktionary-db/getdata"
	"github.com/ninetypercentlanguage/wiktionary-db/savedata"

	_ "github.com/lib/pq" // this is how database/sql recognizes postgres driver
	"github.com/ninetypercentlanguage/wiktionary-db/dbsetup"
	"github.com/ninetypercentlanguage/wiktionary-db/getflags"
)

func main() {
	flagVals := getflags.GetFlags()
	db, err := sql.Open("postgres", flagVals.Db)
	checkErr(err)
	err = dbsetup.Drop(db)
	checkErr(err)
	err = dbsetup.Create(db)
	checkErr(err)

	data := getdata.GetWordsData(flagVals.Data)
	savedata.Save(data, db)
}

func checkErr(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
