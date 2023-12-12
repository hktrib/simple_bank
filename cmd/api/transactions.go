package api

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type PostTransactionRequest struct {
	Email  string `json:"user_email"`
	Date   string `json:"date_of_transaction"`
	Amount string `json:"amount"`
}

func (srv *Server) PostTransactionHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := PostTransactionRequest{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Debug().Err(err).Msg("failed to decode PostTransaction request")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if requestBody.Email != "" && requestBody.Date != "" && requestBody.Amount != "" {

	}

}

func (srv *Server) GetTransactionPDFHandler(w http.ResponseWriter, r *http.Request) {

}
