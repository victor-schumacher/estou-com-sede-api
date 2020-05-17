package plant

import (
	"encoding/json"
	"fmt"
	"github.com/victor-schumacher/estou-com-sede-api/server"
	"net/http"
)

type Manager struct {
}

func NewPlantManager() Manager {
	return Manager{}
}

func (m Manager) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			processPost(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func processPost(w http.ResponseWriter, r *http.Request) {
	plant := server.Plant{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&plant)
	if err != nil {
		http.Error(w, "Error while processing data", http.StatusBadRequest)
	}
	humidityLevel := humidityLevel(plant.SensorHumidity)
	fmt.Println(humidityLevel)
}

// Translating sensor humidity to four levels
func humidityLevel(sh int) int {
	switch {
	case sh >= 100:
		return 4
	case sh > 50 && sh < 75:
		return 3
	case sh > 25 && sh < 50:
		return 2
	case sh < 25:
		return 1
	default:
		return 0
	}
}
