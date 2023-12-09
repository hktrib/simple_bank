package db

import (
	"errors"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog/log"
)

var (
	DatabaseFile = "database.csv"
	layoutString = "12/12/2020"
)

type Database struct {
	Records []transaction
}

type transaction struct {
	Email  string `csv:"user_email"`
	Date   string `csv:"date_of_transaction"`
	Amount string `csv:"amount"`
}

// FilterRecords reads from DatabaseFile and stores filter records based on query
func (db *Database) FilterRecords(startDate time.Time, endDate time.Time) error {
	file, err := os.Open(DatabaseFile)
	if err != nil {
		log.Debug().Err(err).Msg("Error opening database file")
		return err
	}

	var rawDB Database

	// Getting the entire 'database' from csv
	if err := gocsv.UnmarshalFile(file, &rawDB.Records); err != nil {
		log.Debug().Err(err).Msg("Unable to Unmarshald database csv file")
		return err
	}

	// Data being in correct format is insured by the API client

	// Filtering DB []Records
	for _, record := range rawDB.Records {
		transactionDate, err := time.Parse(layoutString, record.Date)
		if err != nil {
			return errors.New("internal server error")
		}

		if startDate.Equal(transactionDate) || endDate.Equal(transactionDate) {
			db.Records = append(db.Records, record)
		} else if startDate.Before(transactionDate) && endDate.After(transactionDate) {
			db.Records = append(db.Records, record)
		}
	}

	defer file.Close()
	return nil
}

// func (db *DB) FilterRecords() error {

// }
