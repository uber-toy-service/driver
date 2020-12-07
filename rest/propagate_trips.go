package rest_api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type PropagationTripBucket struct {
	TripId    string   `json:"trip_id"`
	DriverIds []string `json:"driver_ids"`
}

func PropagateTripsFromUpstream(w http.ResponseWriter, req *http.Request) {
	var tripBucket PropagationTripBucket
	json.NewDecoder(req.Body).Decode(&tripBucket)
	for _, val := range tripBucket.DriverIds {
		data := map[string]string{
			"trip_id": tripBucket.TripId,
			"driver":  val,
		}
		pusherClient.Trigger("trip-channel", "trip-appeared", data)
	}
}

func InitTripPropagation(r *mux.Router) {
	r.HandleFunc("/api/driver/trip_broadcast", PropagateTripsFromUpstream).Methods("POST").Headers("Content-Type", "application/json")
}
