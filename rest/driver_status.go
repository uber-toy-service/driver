package rest_api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	dbHost = "https://uber-main-db.herokuapp.com"
)

type DriverStatus struct {
	Id             string  `json:"id"`
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
	// this is not included in POST
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

// Accepts Driver by id and checks its status using the database,
// returns JSON-encoded result
func GetDriverStatus(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var status DriverStatus
	resp, _ := http.Get(dbHost + "/api/front/car/" + vars["driver_id"])
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&status)
}

// Writes Driver status to the database
func PostDriverStatus(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp, _ := http.Post(dbHost+"/api/front/car/"+vars["driver_id"], "application/json", req.Body)
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := ioutil.ReadAll(resp.Body)
	w.Write(bytes)
}
