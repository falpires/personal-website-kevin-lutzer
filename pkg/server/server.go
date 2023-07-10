package server

import (
	"net/http"

	"github.com/kevinlutzer/personal-website-api/pkg/visitor"
)

type server struct {
	visitorService visitor.Service
}

type Server interface {
	CreateVisitor(w http.ResponseWriter, r *http.Request)
	ListVisitor(w http.ResponseWriter, r *http.Request)
}

func NewServer(visitorService visitor.Service) Server {
	return &server{
		visitorService: visitorService,
	}
}
