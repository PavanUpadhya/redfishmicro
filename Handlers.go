package main

import(
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
)

//ServerDetails struct to hold server IP and credentials
type ServerDetails struct{
	IPAddress string
	UserName string
	Password string
}

func discoverBasicInfoHandler(w http.ResponseWriter, r *http.Request) {
	var serverdetails = new(ServerDetails)
	err := json.NewDecoder(r.Body).Decode(&serverdetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !isRequestValid(serverdetails){
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Server Details Json : %+v", serverdetails)
	w.WriteHeader(http.StatusOK)
}

func discoverPowerState(w http.ResponseWriter, r *http.Request) {
	var serverdetails = new(ServerDetails)
	err := json.NewDecoder(r.Body).Decode(&serverdetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !isRequestValid(serverdetails){
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Server Details Json : %+v", serverdetails)
	w.WriteHeader(http.StatusOK)
}

func isRequestValid(servDetails *ServerDetails) bool {
	if len(strings.Trim(servDetails.IPAddress, " ")) == 0 ||
		len(strings.Trim(servDetails.UserName, " ")) == 0 ||
		len(strings.Trim(servDetails.Password, " ")) == 0 {
			return false
	}
	return true
}