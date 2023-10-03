package router

import (
	"net/http"

	"github.com/ImranZahoor/blog-api/internal/controller"
	"github.com/gorilla/mux"
)

type (
	Server struct {
		controller controller.Controller
		router     *mux.Router
	}
)

func NewServer() *Server {
	return &Server{router: mux.NewRouter()}
}

func (s *Server) RegisterHandlers() {
	articleRouter := s.router.PathPrefix("/article").Subrouter()
	articleRouter.HandleFunc("/", s.controller.ListArticleHandler).Methods(http.MethodGet)
	articleRouter.HandleFunc("/{id}", s.controller.GetArticleByIDHandler).Methods(http.MethodGet)
	articleRouter.HandleFunc("/", s.controller.CreateArticleHandler).Methods(http.MethodPost)
	articleRouter.HandleFunc("/{id}", s.controller.UpdateArticleHandler).Methods(http.MethodPut)
	articleRouter.HandleFunc("/{id}", s.controller.DeleteArticleHandler).Methods(http.MethodDelete)
}

func (s *Server) GetRouter() *mux.Router {
	return s.router
}
