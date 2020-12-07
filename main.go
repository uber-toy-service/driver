package main

import (
	"net/http"
	"os"

	rest_api "driver/rest"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	rest_api.InitBroadcastHandling(router)
	rest_api.InitDriverStatus(router)
	rest_api.InitTripPropagation(router)
	rest_api.InitDriverAcceptsTrip(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
