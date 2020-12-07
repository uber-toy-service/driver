package rest_api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	SupplyLocaiton string = "https://toy-supply-location-service.herokuapp.com"
)

type LocationMsg struct {
	Id       string   `json:"id"`
	Location Location `json:"location"`
}

type Location struct {
	Longtitude float64 `json:"longtitude"`
	Latitude   float64 `json:"latitude"`
}

func PassDriversLocationToSupply(msg LocationMsg) []byte {
	msgJSON, _ := json.Marshal(&msg)
	buffer := bytes.NewBuffer(msgJSON)
	resp, _ := http.Post(SupplyLocaiton+"/update-supply", "application/json", buffer)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func PassDriversLocationToDB(driver_id string, msg LocationMsg) []byte {
	msgJSON, _ := json.Marshal(&msg)
	buffer := bytes.NewBuffer(msgJSON)
	req, _ := http.NewRequest("PATCH", dbHost+"/api/front/car/"+driver_id, buffer)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// Initiates updates of Driver's location
func UpdateDriversLocation(w http.ResponseWriter, req *http.Request) {
	driver_id := mux.Vars(req)["driver_id"]

	var location LocationMsg
	json.NewDecoder(req.Body).Decode(&location)

	locationMsg := LocationMsg{Id: driver_id, Location: location.Location}
	_ = PassDriversLocationToSupply(locationMsg)
	_ = PassDriversLocationToDB(driver_id, locationMsg)

	// Respond to Rostyk
}

func InitBroadcastHandling(r *mux.Router) {
	r.HandleFunc("/api/driver/{driver_id}", UpdateDriversLocation).Methods("POST")
}
