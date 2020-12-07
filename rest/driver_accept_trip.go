package rest_api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	tripCache map[string]string
	mutex     *sync.Mutex
)

type SuccessStatus struct {
	Status string `json:"status"`
}

type AcceptTripMsg struct {
	DriverId string `json:"driver_id"`
	TripId   string `json:"trip_id"`
}

func AcceptTrip(w http.ResponseWriter, req *http.Request) {
	driverId := mux.Vars(req)["driver_id"]
	var msg AcceptTripMsg
	json.NewDecoder(req.Body).Decode(&msg)

	var status SuccessStatus

	mutex.Lock()
	if _, ok := tripCache[msg.TripId]; !ok {
		tripCache[msg.TripId] = driverId
		status = SuccessStatus{Status: "ok; your response has been recorded"}
	} else {
		status = SuccessStatus{Status: "bad; this trip is already taken"}
	}
	mutex.Unlock()

	json.NewEncoder(w).Encode(&status)
}

func InitDriverAcceptsTrip(r *mux.Router) {
	mutex = &sync.Mutex{}
	r.HandleFunc("/api/driver/accept_trip{driver_id}", AcceptTrip).Methods("POST").Headers("Content-Type", "application/json")
}
