package handler

import (
	"encoding/json"
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
	switch req.Method {
	case "GET":
		h.getProblems(w, req)
	case "POST":
		h.createProblem(w, req)
	}
}

func (h *Handler) getProblems(w http.ResponseWriter, req *http.Request) {
	problems, err := h.service.GetProblems()
	if err != nil {
		http.Error(w, "failed to get list of problems", http.StatusInternalServerError)
		return
	}

	bs, err := json.Marshal(map[string]interface{}{
		"problems": problems,
	})
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(bs)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) createProblem(w http.ResponseWriter, req *http.Request) {
	input := map[string]string{}

	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "failed to read statement", http.StatusBadRequest)
		return
	}

	if _, ok := input["statement"]; !ok {
		http.Error(w, "missing 'statement' field in json payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateProblem(input["statement"])
	if err != nil {
		http.Error(w, "problem creating problem", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
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
	switch req.Method {
	case "GET":
	case "POST":
	}
}
