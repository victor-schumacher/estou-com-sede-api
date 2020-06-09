package handler

import (
	"encoding/json"
	"github.com/victor-schumacher/estou-com-sede-api/api/twitter"
	"io/ioutil"
	"net/http"
	"strconv"

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

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading", http.StatusBadRequest)
	}

	err = json.Unmarshal(bodyBytes, &plant)
	if err != nil {
		http.Error(w, "Error unmarshal", http.StatusBadRequest)
	}

	hInt, err := strconv.Atoi(plant.SensorHumidity)

	humidityLevel := service.HumidityLevel(hInt)

	twitter.PostTweet(humidityLevel)

	//email.SendEmail()
}
