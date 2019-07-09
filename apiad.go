package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var p = fmt.Printf
var l = log.Fatalf

func createUser(w http.ResponseWriter, r *http.Request) {

}

func disableUser(w http.ResponseWriter, r *http.Request) {

}

func enableUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func getLockedUsers(w http.ResponseWriter, r *http.Request) {

}

func unlockUser(w http.ResponseWriter, r *http.Request) {

}

func getUserAttribute(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//If no arguments were specified, the length of os.Args will be 1
	port := ""
	if len(os.Args) < 2 {
		port = "8000"
	} else {
		port = os.Args[1]
	}

	r := mux.NewRouter()

	r.HandleFunc("/v1/adduser", createUser).Methods("POST")
	r.HandleFunc("/v1/lockedusers", getLockedUsers).Methods("GET")
	r.HandleFunc("/v1/unlock/{sam}", unlockUser).Methods("POST")
	r.HandleFunc("/v1/disable/{sam}", disableUser).Methods("POST")

	p("listening port %s...\n", port)
	port = fmt.Sprintf(":%s", port)
	http.ListenAndServe(port, r)

}
