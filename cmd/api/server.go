package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	db "github.com/hktrib/simple_bank/internal/database"
	pdf "github.com/hktrib/simple_bank/internal/pdf"
)

type Server struct {
	Router       chi.Router
	DB           *db.Database
	PDFGenerator *pdf.PDFGenerator
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
		DB:     &db.Database{},
	}
}

func (srv *Server) MountHandlers() {
	srv.Router.Use(middleware.Logger)
	srv.Router.Post("/addtransaction", srv.PostTransactionHandler)
	srv.Router.Post("/filtertransactions", srv.GetTransactionPDFHandler)
}
