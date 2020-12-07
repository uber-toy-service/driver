package rest_api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

type AcceptTripMsg struct {
	DriverId string `json:"driver_id"`
	TripId   string `json:"trip_id"`
}

// TODO StartTrip, EndTrip, CancelTrip

func StartTrip(w http.ResponseWriter, req *http.Request) {
	var msg AcceptTripMsg
	json.NewDecoder(req.Body).Decode(&msg)

	data := url.Values{}
	data.Set("driver_id", msg.DriverId)
	data.Set("trip_id", msg.TripId)

	resp, _ := http.Post(SupplyLocaiton+"/driver/accept-trip", "application/json", strings.NewReader(data.Encode()))
	defer resp.Body.Close()

	w.Write(resp)
}

func AcceptTrip(w http.ResponseWriter, req *http.Request) {
	var msg AcceptTripMsg
	json.NewDecoder(req.Body).Decode(&msg)

	data := url.Values{}
	data.Set("driver_id", msg.DriverId)
	data.Set("trip_id", msg.TripId)

	http.Post(SupplyLocaiton+"/driver/accept-trip", "application/json", strings.NewReader(data.Encode()))
}

func InitDriverAcceptsTrip(r *mux.Router) {
	r.HandleFunc("/api/driver/accept_trip", AcceptTrip).Methods("POST").Headers("Content-Type", "application/json")
}
