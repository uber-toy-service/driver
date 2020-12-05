package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Post that handles Gleb's upstream

type DriverStatus struct {
	Id             int
	Email          string
	Phone          string
	FirstName      string
	MiddleName     string
	LastName       string
	Capacity       int
	CarClassId     int
	Note           string
	AcceptsRides   bool
	OnTheRide      bool
	CarStatusId    int
	CoordLatitude  float64
	CoordLongitude float64
}

// Accepts Driver by id and checks its status using the database,
// returns the JSON-encoded result
func GetDriverStatus(w http.ResponseWriter, req *http.Request) {

	// use the DB instead
	status := DriverStatus{id: 0}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&status)
}

/* This is done using the Pusher API
func DriverAcceptTrip(w http.ResponseWriter, req *http.Request) {
	// driver_id
	// read json
        }
*/

// Updates the location of the Driver
func UpdateDriversLocation(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/driver/{driver_id}", GetDriverStatus).Methods("GET")
	// router.HandleFunc("/api/driver", DriverAcceptTrip).Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)

}
