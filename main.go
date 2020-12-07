package main

import (
	"net/http"
	"os"

	rest_api "driver/rest"

	"github.com/gorilla/mux"
)

/* This is done using the Pusher API
func DriverAcceptTrip(w http.ResponseWriter, req *http.Request) {
	// driver_id
	// read json
        }
*/

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/driver/{driver_id}", rest_api.GetDriverStatus).Methods("GET")
	router.HandleFunc("/api/driver/{driver_id}", rest_api.PostDriverStatus).Methods("POST")
	router.HandleFunc("/api/driver/{driver_id}", rest_api.UpdateDriversLocation).
		Methods("POST")
	rest_api.InitTripPropagation(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
