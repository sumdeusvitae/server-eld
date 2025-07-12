package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Driver struct moved here — only define once
type Driver struct {
	Name             string    `bson:"name" json:"name"`
	Status           string    `bson:"status" json:"status"`
	Location         string    `bson:"location" json:"location"`
	TruckID          string    `bson:"truck_id" json:"truck_id"`
	ShiftStart       string    `bson:"shift_start" json:"shift_start"`
	BreakTime        string    `bson:"break_time" json:"break_time"`
	DriveTime        string    `bson:"drive_time" json:"drive_time"`
	CycleTime        string    `bson:"cycle_time" json:"cycle_time"`
	ConnectionStatus string    `bson:"connection_status" json:"connection_status"`
	ReportedAt       string    `bson:"reported_at" json:"reported_at"`
	LastUpdated      time.Time `bson:"last_updated" json:"last_updated"`
}

// Shared state
var (
	cachedDrivers []Driver
	lastFetched   time.Time
)

// Exported handler
func DriversHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleDriversGET(w, r)
	case http.MethodPost:
		handleDriversPOST(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func handleDriversGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(cachedDrivers) == 0 {
		log.Println("No drivers cached yet")
	}

	json.NewEncoder(w).Encode(cachedDrivers)
}

func handleDriversPOST(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Drivers []Driver `json:"drivers"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if len(req.Drivers) == 0 {
		http.Error(w, "No drivers provided", http.StatusBadRequest)
		return
	}

	for _, d := range req.Drivers {
		d.LastUpdated = time.Now()
		cachedDrivers = append(cachedDrivers, d)
	}

	log.Printf("Received %d drivers\n", len(req.Drivers))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}
