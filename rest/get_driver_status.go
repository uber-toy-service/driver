package rest_api

import (
	"encoding/json"
	"net/http"
)

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
