package rest_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"driver/logging"

	"github.com/gorilla/mux"
)

const (
	SupplyLocaiton string = "https://toy-supply-location-service.herokuapp.com"
)

// DONE

type LocationWithIdMsg struct {
	id       string   `json:"id"`
	location Location `json:"location"`
}

type LocationMsg struct {
	Location Location `json:"location"`
}

type Location struct {
	Longtitude float64 `json:"longtitude"`
	Latitude   float64 `json:"latitude"`
}

type UpdateDriversLocationResponse struct {
	Status string `json:"status"`
}

func PassDriversLocationToSupply(msg LocationWithIdMsg) []byte {
	msgJSON, _ := json.Marshal(msg)
	buffer := bytes.NewBuffer(msgJSON)
	req, _ := http.NewRequest("POST", SupplyLocaiton+"/update-supply", buffer)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// TODO Work with Yegor
func StoreDriversLocationInDB(msg LocationWithIdMsg) []byte {
	return []byte(`foo`)
}

// Initiates updates of Driver's location
func UpdateDriversLocation(w http.ResponseWriter, req *http.Request) {
	driver_id := mux.Vars(req)["driver_id"]
	logging.DebugFile.Println("Driver's id=", driver_id)
	var location LocationMsg
	err := json.NewDecoder(req.Body).Decode(&location)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Cannot update the location of the driver due to invalid JSON passed", http.StatusUnprocessableEntity)
		return
	}

	locationWithIdMsg := LocationWithIdMsg{id: driver_id, location: location.Location}
	_ = PassDriversLocationToSupply(locationWithIdMsg)

	// Send to DB

	// Respond to Rostyk
	status := UpdateDriversLocationResponse{Status: "ok"}
	json.NewEncoder(w).Encode(status)
}
