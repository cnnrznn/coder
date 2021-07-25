package main

import (
	"fmt"
	"net/http"

	"github.com/cnnrznn/coder/handler"
	"github.com/cnnrznn/coder/repository"
	"github.com/cnnrznn/coder/service"
	log "github.com/golang/glog"
	"github.com/gorilla/mux"
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

	r.HandleFunc("/problems", h.HandleProblems).Methods("GET", "POST")                                    // list problems, create a problem
	r.HandleFunc("/problems/{pid}", h.HandleProblem).Methods("GET")                                       // get a specific problem
	r.HandleFunc("/problems/{pid}/testcases", h.HandleTestCases).Methods("GET", "POST")                   // list test cases, create test case
	r.HandleFunc("/problems/{pid}/testcases/{tid}/submit", h.HandleTestCaseSubmit).Methods("GET", "POST") // get input for test case, submit response

	log.Info("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
