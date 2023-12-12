package main

import (
	"net/http"

	"github.com/hktrib/simple_bank/cmd/api"
	"github.com/hktrib/simple_bank/internal/db"
)

func main() {

	srv := api.NewServer()

	csvDB := db.Database{}

	// // Temp Start/End times

	// startDate, err := time.Parse(db.LayoutString, "12/02/2023")
	// if err != nil {
	// 	log.Debug().Err(err).Msg("error parsing time")
	// }
	// endDate, err := time.Parse(db.LayoutString, "12/08/2023")
	// if err != nil {
	// 	log.Debug().Err(err).Msg("error parsing time")
	// }

	// if startDate.IsZero() || endDate.IsZero() {
	// 	fmt.Println("Invalid date format")
	// 	return
	// }
	// err = csvDB.FilterRecords("hktribunal@gmail.com", startDate, endDate)
	// if err != nil {
	// 	fmt.Println("Something went wrong")
	// 	log.Debug().Err(err).Msg("erorr")
	// 	return
	// }
	// fmt.Println(csvDB.FilteredRecords)

	// mux := chi.NewRouter()

	go func() {

		// Starting the server
		http.ListenAndServe(":8080", srv.Router)
	}()
	select {}
}
