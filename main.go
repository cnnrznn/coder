package main

import (
	"fmt"
	"net/http"

	log "github.com/golang/glog"
	"github.com/gorilla/mux"

	"github.com/cnnrznn/coder/handler"
	"github.com/cnnrznn/coder/repository"
	"github.com/cnnrznn/coder/service"
)

func main() {
	fmt.Println("Welcome to CODER ^_^")

	r := mux.NewRouter()

	repo := repository.New()

	err := repo.Connect()
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}

	s := service.New(repo)
	h := handler.New(s)

	r.HandleFunc("/problems", h.HandleProblems).Methods("GET")
	r.HandleFunc("/problems/{pid}/statement", h.HandleProblemStatement).Methods("GET")
	r.HandleFunc("/problems/{pid}/testcases/{tid}", h.HandleTestCase).Methods("GET", "POST")

	log.Info("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
