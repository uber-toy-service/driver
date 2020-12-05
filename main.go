package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// TODO Use real data, for now, mock everything

// Post that handles Gleb's upstream

type DriverStatus struct {
	Id             int     `json:"id"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	FirstName      string  `json:"first_name"`
	MiddleName     string  `json:"middle_name"`
	LastName       string  `json:"last_name"`
	Capacity       int     `json:"capacity"`
	CarClassId     int     `json:"car_class_id"`
	Note           string  `json:"note"`
	AcceptsRides   bool    `json:"accepts_rides"`
	OnTheRide      bool    `json:"on_the_ride"`
	CarStatusId    int     `json:"car_status_id"`
	CoordLatitude  float64 `json:"coord_latitude"`
	CoordLongitude float64 `json:"coord_longtitude"`
}

func getDBJSONResult() {}

// Accepts Driver by id and checks its status using the database,
// returns the JSON-encoded result
func GetDriverStatus(w http.ResponseWriter, req *http.Request) {
	// use the DB instead
	status := DriverStatus{Id: 0}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&status)
}

/* This is done using the Pusher API
func DriverAcceptTrip(w http.ResponseWriter, req *http.Request) {
	// driver_id
	// read json
        }
*/

type Location struct {
	Longtitude float64 `json:"longtitude"`
	Latitude   float64 `json:"latitude"`
}

// Updates the location of the Driver
func UpdateDriversLocation(w http.ResponseWriter, req *http.Request) {
	var location Location
	err := json.NewDecoder(req.Body).Decode(&location)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Received data", req.Body)
	fmt.Println("Received the location", location)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/driver/{driver_id}", GetDriverStatus).Methods("GET").Headers("Content-Type", "application/json")
	router.HandleFunc("/api/driver/{driver_id}", UpdateDriversLocation).
		Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
