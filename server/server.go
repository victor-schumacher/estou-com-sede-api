package server

import "net/http"

type  Plant struct {
	SensorHumidity int `json:"sensorHumidity"`
	FinalHumidity int `json:"finalHumidity"`
}

type PlantHandler interface {
	Handle() http.HandlerFunc
}