package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	ErrFile   *log.Logger
	DebugFile *log.Logger
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

/// Business logic

// TODO Post that handles Gleb's upstream

// TODO Use real data, for now, mock everything

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

// This talks to the database
func getDBJSONDriverStatus(res *DriverStatus) {
	res.Id = 1
}

// Accepts Driver by id and checks its status using the database,
// returns JSON-encoded result
func GetDriverStatus(w http.ResponseWriter, req *http.Request) {
	var status DriverStatus
	getDBJSONDriverStatus(&status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&status)
}

// DONE

type Location struct {
	Longtitude float64 `json:"longtitude"`
	Latitude   float64 `json:"latitude"`
}

type UpdateDriversLocationResponse struct {
	Status string `json:"status"`
}

// Updates Driver's location in the database
func UpdateDriversLocation(w http.ResponseWriter, req *http.Request) {
	driver_id := mux.Vars(req)["driver_id"]
	DebugFile.Println("Driver's id=", driver_id)
	var location Location
	err := json.NewDecoder(req.Body).Decode(&location)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Cannot update the location of the driver due to invalid JSON passed", http.StatusUnprocessableEntity)
		return
	}
	/*
		location.Latitude += 1
		location.Longtitude += 1
	*/
	status := UpdateDriversLocationResponse{Status: "ok"}
	json.NewEncoder(w).Encode(status)
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
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
