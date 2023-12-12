package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	db "github.com/hktrib/simple_bank/internal/database"
)

type PostTransactionRequest struct {
	Email  string `json:"user_email"`
	Date   string `json:"date_of_transaction"`
	Amount string `json:"amount"`
}

type FilterTransactionsRequest struct {
	Email     string `json:"user_email"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (srv *Server) PostTransactionHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := PostTransactionRequest{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Debug().Err(err).Msg("failed to decode /addtransaction request")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestBody.Email == "" && requestBody.Date == "" && requestBody.Amount == "" {

		log.Debug().Err(err).Msg("request doesn't have all params")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = srv.DB.AddRecord(&db.Transaction{
		Email:  requestBody.Email,
		Date:   requestBody.Date,
		Amount: requestBody.Amount,
	})
	if err != nil {
		log.Debug().Msg("failed in adding transaction to database")
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Debug().Msg("succeeded in adding transaction to database")

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK!"))
}

func (srv *Server) GetTransactionPDFHandler(w http.ResponseWriter, r *http.Request) {

	requestBody := FilterTransactionsRequest{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Debug().Err(err).Msg("failed to decode /filtertransactions request")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestBody.Email == "" && requestBody.StartDate == "" && requestBody.EndDate == "" {
		log.Debug().Err(err).Msg("request doesn't have all params")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse(db.LayoutString, requestBody.StartDate)
	if err != nil {
		log.Debug().Err(err).Msg("error parsing time")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse(db.LayoutString, requestBody.EndDate)
	if err != nil {
		log.Debug().Err(err).Msg("error parsing time")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if startDate.IsZero() || endDate.IsZero() {
		fmt.Println("Invalid date format")
		return
	}
	err = srv.DB.FilterRecords(requestBody.Email, startDate, endDate)
	if err != nil {
		fmt.Println("Something went wrong")
		log.Debug().Err(err).Msg("erorr")
		return
	}

	fmt.Println(srv.DB.FilteredRecords)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK!"))

}
