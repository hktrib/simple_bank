package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// fmt.Println("Starting")

	// file, err := os.Open("database.csv")
	// if err != nil {
	// 	log.Debug().Err(err).Msg("Error opening database file")
	// 	return
	// }

	// db := db.Database{}

	// reader := csv.NewReader(file)
	// reader.FieldsPerRecord = -1 // Allow variable number of fields
	// data, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }

	// // Print the CSV data
	// for _, row := range data {
	// 	for _, col := range row {
	// 		fmt.Printf("%s,", col)
	// 	}
	// 	fmt.Println()
	// }

	mux := chi.NewRouter()

	go func() {

		// Starting the server
		http.ListenAndServe(":8080", mux)
	}()
	select {}
}
