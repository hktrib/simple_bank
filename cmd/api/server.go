package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	db "github.com/hktrib/simple_bank/internal/database"
	"github.com/hktrib/simple_bank/internal/email"
	pdf "github.com/hktrib/simple_bank/internal/pdf"
)

type Server struct {
	Router       chi.Router
	DB           *db.Database
	PDFGenerator *pdf.PDFGenerator
	Emailer      *email.Emailer
}

func NewServer(senderEmail string, senderPasscode string) *Server {
	return &Server{
		Router:       chi.NewRouter(),
		DB:           &db.Database{},
		PDFGenerator: &pdf.PDFGenerator{},
		Emailer: &email.Emailer{
			Sender:   senderEmail,
			Passcode: senderPasscode,
		},
	}
}

func (srv *Server) MountHandlers() {
	srv.Router.Use(middleware.Logger)
	srv.Router.Post("/addtransaction", srv.PostTransactionHandler)
	srv.Router.Post("/filtertransactions", srv.GetTransactionPDFHandler)
}
