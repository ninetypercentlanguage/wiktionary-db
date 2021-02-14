package getflags

import (
	"flag"
	"fmt"
)

type flagValues struct {
	Data string // path to combined data directory
	Db   string // postgres connection string
}

func GetFlags() flagValues {
	dataPathPtr := flag.String("data", "", "the path to the output of wiktionary-combine")
	dbConnectionPtr := flag.String("db", "", "a connection string for postgres db")

	flag.Parse()

	fv := flagValues{
		Data: *dataPathPtr,
		Db:   *dbConnectionPtr,
	}
	if fv.Data == "" || fv.Db == "" {
		fmt.Println("Must provide the following flags:")
		flag.PrintDefaults()
		panic("missing flags")
	}
	return fv
}
