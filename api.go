package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APISrver struct {
	listenAddr string
	store      storage
}

func NewAPIServer(listenAddr string, store storage) *APISrver {
	return &APISrver{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APISrver) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APISrver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("methos not allowed %s", r.Method)
}

func (s *APISrver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	id := mux.Vars(r)["id"]
	fmt.Println(id)
	//db.get(id)

	return WriteJSON(w, http.StatusOK, &Account{})

}
func (s *APISrver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APISrver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APISrver) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func WriteJSON(w http.ResponseWriter, status int, v any) error {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
