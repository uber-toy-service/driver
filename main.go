package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Post that handles Gleb's upstream

// Accepts Driver by id and checks its status using the database,
// returns the JSON-encoded result
func DriverGetStatus(w http.ResponseWriter, req *http.Request) {
	foo := map[string]string{"name": "foo"}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(foo)
	json.NewEncoder(w).Encode(&foo)
}

/* This is done using the Pusher API
func DriverAcceptTrip(w http.ResponseWriter, req *http.Request) {
	// driver_id
	// read json
        }
*/

// Updates the location of the Driver
func UpdateDriversLocation(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/driver", DriverStatus).Methods("GET")
	router.HandleFunc("/api/driver", DriverAcceptTrip).Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)

}
