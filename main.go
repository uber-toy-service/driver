package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func DriverGetStatus(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "GET request")
}

func DriverAcceptTrip(w http.ResponseWriter, req *http.Request) {
	//
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/driver", DriverGetStatus).Methods("GET")
	router.HandleFunc("/api/driver", DriverAcceptTrip).Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
