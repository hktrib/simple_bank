package db

import (
	"encoding/csv"
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
	RawRecords               []Transaction
	FilteredRecords          []Transaction
	FilteredParseableRecords [][]string
}

type Transaction struct {
	Email  string `csv:"user_email"`
	Date   string `csv:"date_of_transaction"`
	Amount string `csv:"amount"`
}

// FilterRecords reads from DatabaseFile and stores filter records based on queried startDate and endDate
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

	// initializing columns for filtered records.
	columns := []string{}
	columns = append(columns, "user_email")
	columns = append(columns, "date_of_transaction")
	columns = append(columns, "amount")
	db.FilteredParseableRecords = append(db.FilteredParseableRecords, columns)

	// Filtering DB []Records
	for _, record := range db.RawRecords {

		filteredTransactionData := []string{}
		transactionDate, err := time.Parse(LayoutString, record.Date)
		if err != nil {
			return errors.New("internal server error")
		}

		if record.Email == userEmail {
			if startDate.Equal(transactionDate) || endDate.Equal(transactionDate) {

				// if transactionDate is equal to either startDate or endDate

				filteredTransactionData = append(filteredTransactionData, record.Email)
				filteredTransactionData = append(filteredTransactionData, record.Date)
				filteredTransactionData = append(filteredTransactionData, record.Amount)

				db.FilteredRecords = append(db.FilteredRecords, record)
				db.FilteredParseableRecords = append(db.FilteredParseableRecords, filteredTransactionData)
			} else if startDate.Before(transactionDate) && endDate.After(transactionDate) {

				// Else If transactionDate is in range between startDate and endDate

				filteredTransactionData = append(filteredTransactionData, record.Email)
				filteredTransactionData = append(filteredTransactionData, record.Date)
				filteredTransactionData = append(filteredTransactionData, record.Amount)

				db.FilteredRecords = append(db.FilteredRecords, record)
				db.FilteredParseableRecords = append(db.FilteredParseableRecords, filteredTransactionData)
			}
		}
	}

	// error: unable to filter returning nil
	return nil
}

// Adds a new Transaction record to the database csv file.
func (db *Database) AddRecord(record *Transaction) error {

	database, err := os.OpenFile(DatabaseFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("unable to open database file")
	}
	defer database.Close()

	if record != nil {
		rawTransactionData := []string{}

		rawTransactionData = append(rawTransactionData, record.Email)
		rawTransactionData = append(rawTransactionData, record.Date)
		rawTransactionData = append(rawTransactionData, record.Amount)

		fmt.Println(rawTransactionData)

		w := csv.NewWriter(database)
		w.Write(rawTransactionData)
		if err := w.Error(); err != nil {
			return fmt.Errorf("unable to write record: %v", err)
		}
		w.Flush()
		return nil
	}
	return errors.New("no transaction record to add")
}

func (db *Database) Clear() {
	db.FilteredParseableRecords = [][]string{}
	db.FilteredRecords = []Transaction{}
	db.RawRecords = []Transaction{}
}
