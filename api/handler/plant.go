package handler

import (
	"encoding/json"
	"github.com/victor-schumacher/estou-com-sede-api/api/twitter"
	"net/http"

	"github.com/victor-schumacher/estou-com-sede-api/pkg/entity"
	"github.com/victor-schumacher/estou-com-sede-api/pkg/service"
)

type Manager struct {
}

type PlantHandler interface {
	Handle() http.HandlerFunc
}

func NewPlantManager() Manager {
	return Manager{}
}

func (m Manager) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			processPost(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func processPost(w http.ResponseWriter, r *http.Request) {
	plant := entity.Plant{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&plant)
	if err != nil {
		http.Error(w, "Error while processing data", http.StatusBadRequest)
	}
	humidityLevel := service.HumidityLevel(plant.SensorHumidity)

	go twitter.PostTweet(humidityLevel)


}
