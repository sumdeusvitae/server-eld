package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Driver struct moved here — only define once
type Driver struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	Name             string             `bson:"name" json:"name"`
	Status           string             `bson:"status" json:"status"`
	Location         string             `bson:"location" json:"location"`
	TruckID          string             `bson:"truck_id" json:"truck_id"`
	ShiftStart       string             `bson:"shift_start" json:"shift_start"`
	BreakTime        string             `bson:"break_time" json:"break_time"`
	DriveTime        string             `bson:"drive_time" json:"drive_time"`
	CycleTime        string             `bson:"cycle_time" json:"cycle_time"`
	ConnectionStatus string             `bson:"connection_status" json:"connection_status"`
	ReportedAt       string             `bson:"reported_at" json:"reported_at"`
	LastUpdated      time.Time          `bson:"last_updated" json:"last_updated"`
}

// Shared state
var (
	cachedDrivers []Driver
	lastFetched   time.Time
	cacheDuration = 10 * time.Minute
)

// Exported handler
func DriversHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		http.Error(w, "MONGO_URI not set", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		http.Error(w, "client couldnt", http.StatusInternalServerError)
		return
	}

	collection = client.Database("eld_data").Collection("drivers")

	if cachedDrivers == nil || now.Sub(lastFetched) > cacheDuration {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			http.Error(w, "Database query failed", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		var drivers []Driver
		for cursor.Next(ctx) {
			var d Driver
			if err := cursor.Decode(&d); err != nil {
				log.Println("Error decoding driver:", err)
				continue
			}
			drivers = append(drivers, d)
		}

		if err := cursor.Err(); err != nil {
			http.Error(w, "Cursor error", http.StatusInternalServerError)
			return
		}

		cachedDrivers = drivers
		lastFetched = now
		log.Println("Drivers refreshed from DB")
	} else {
		log.Println("Serving drivers from cache")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cachedDrivers)
}
