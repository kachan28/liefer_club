package app

import (
	"encoding/csv"
	"log"
	"os"
)

const dbsFilePath = "dbs.csv"

type databaseList struct {
	Dbs []database
}

type database struct {
	UpdateDt string
	Name     string
}

func (d *databaseList) getConf() {
	f, err := os.Open(dbsFilePath)
	if err != nil {
		log.Fatal("can't open updated databases list", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("can't parse file as CSV for "+dbsFilePath, err)
	}
	for _, record := range records {
		d.Dbs = append(d.Dbs, database{UpdateDt: record[0], Name: record[1]})
	}
}
