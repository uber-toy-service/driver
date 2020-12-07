package rest_api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO Discuss trip cancellation

type PropagationTripBucket struct {
	TripId    string `json:"trip_id"`
	DriverIds string `json:"driver_ids"`
}

func PropagateTripsFromUpstream(w http.ResponseWriter, req *http.Request) {
	var tripBucket PropagationTripBucket
	json.NewDecoder(req.Body).Decode(&tripBucket)
	// Send to the drivers
	// Need the Pusher API here.
}

func InitTripPropagation(r *mux.Router) {
	r.HandleFunc("/api/driver/trip_broadcast", PropagateTripsFromUpstream).Methods("GET").Headers("Content-Type", "application/json")
}
