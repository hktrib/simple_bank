package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router chi.Router
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (srv *Server) ServeHTTP() {
	srv.Router.Use(middleware.Logger)
	srv.Router.Post("/addtransaction", srv.PostTransactionHandler)
	srv.Router.Get("/filtertransactions", srv.GeneratePDF)
}
