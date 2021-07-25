package handler

import (
	"net/http"

	"github.com/cnnrznn/coder/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HandleProblems(w http.ResponseWriter, req *http.Request) {
}

func (h *Handler) HandleProblem(w http.ResponseWriter, req *http.Request) {
}

func (h *Handler) HandleTestCaseSubmit(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
	case "POST":
	}
}

func (h *Handler) HandleTestCases(w http.ResponseWriter, req *http.Request) {
}
