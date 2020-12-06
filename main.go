package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	ErrFile   *log.Logger
	DebugFile *log.Logger
)

const (
	SupplyLocaiton string = "https://toy-supply-location-service.herokuapp.com"
)

func init() {
	var errsLogFile, debugLogFile *os.File
	var err error
	if errsLogFile, err = os.OpenFile("errs.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm); err != nil {
		log.Fatalln("Init failed, reason:", err)
	}
	if debugLogFile, err = os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm); err != nil {
		log.Fatalln("Init failed, reason:", err)
	}
	ErrFile = log.New(errsLogFile, "ERROR", log.LstdFlags)
	DebugFile = log.New(debugLogFile, "DEBUG", log.LstdFlags)
}

/* This is done using the Pusher API
func DriverAcceptTrip(w http.ResponseWriter, req *http.Request) {
	// driver_id
	// read json
        }
*/

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/driver/{driver_id}", GetDriverStatus).Methods("GET").Headers("Content-Type", "application/json")
	router.HandleFunc("/api/driver/{driver_id}", UpdateDriversLocation).
		Methods("POST")
	router.HandleFunc("/api/driver/trip_broadcast", PropagateTripsFromUpstream).Methods("GET").Headers("Content-Type", "application/json")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
