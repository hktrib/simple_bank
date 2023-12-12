package db

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog/log"
)

var (
	DatabaseFile = "database.csv"
	LayoutString = "01/02/2006"
)

type Database struct {
	Indexes         []string
	RawRecords      []Transaction
	FilteredRecords []Transaction
}

type Transaction struct {
	Email  string `csv:"user_email"`
	Date   string `csv:"date_of_transaction"`
	Amount string `csv:"amount"`
}

// FilterRecords reads from DatabaseFile and stores filter records based on query
func (db *Database) FilterRecords(userEmail string, startDate time.Time, endDate time.Time) error {
	file, err := os.Open(DatabaseFile)
	if err != nil {
		log.Debug().Err(err).Msg("Error opening database file")
		return err
	}
	defer file.Close()

	// Getting the entire 'database' from csv
	if err := gocsv.UnmarshalFile(file, &db.RawRecords); err != nil {
		log.Debug().Err(err).Msg("Unable to Unmarshald database csv file")
		return err
	}

	// Data being in correct format is insured by the API client

	// Filtering DB []Records
	for _, record := range db.RawRecords {
		transactionDate, err := time.Parse(LayoutString, record.Date)
		if err != nil {
			fmt.Printf("Record.Date: %v\n", record.Date)
			fmt.Printf("LayoutString: %v\n", LayoutString)
			return errors.New("internal server error")
		}

		if record.Email == userEmail {
			if startDate.Equal(transactionDate) || endDate.Equal(transactionDate) {
				db.FilteredRecords = append(db.FilteredRecords, record)
				fmt.Printf("Record: %v\n", record)
			} else if startDate.Before(transactionDate) && endDate.After(transactionDate) {
				db.FilteredRecords = append(db.FilteredRecords, record)
				fmt.Printf("Record: %v\n", record)
				fmt.Printf("Start date: %v\n", startDate)
			} else {
				fmt.Printf("End date: %v\n", endDate)
			}
		}
	}

	return nil
}

func (db *Database) AddRecord(record *Transaction) error {

	database, err := os.OpenFile(DatabaseFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		return errors.New("unable to open database file")
	}

	if record != nil {
		gocsv.Marshal(record, database)
		return nil
	}
	return errors.New("no transaction record to add")
}

// func (db *DB) FilterRecords() error {

// }
