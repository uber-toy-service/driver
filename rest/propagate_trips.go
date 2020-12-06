package rest_api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TODO Discuss trip cancellation
// TODO Post that handles Gleb's upstream

type PropagationTripBucket struct {
	TripId    string `json:"trip_id"`
	DriverIds string `json:"driver_ids"`
}

func PropagateTripsFromUpstream(w http.ResponseWriter, req *http.Request) {
	// var tripBucket PropagationTripBucket
	// res := json.Unmarshal(req.Body, &tripBucket)

	// json.Unmarshal(bytes.NewBuffer(req.Body), &tripBucket)

}

func InitTripPropagation(r *mux.Router) {
	r.HandleFunc("/api/driver/trip_broadcast", PropagateTripsFromUpstream).Methods("GET").Headers("Content-Type", "application/json")
}
