package service

import "github.com/victor-schumacher/estou-com-sede-api/pkg/entity"

func HumidityLevel(sensorHumidity int) string {
	switch {
	case sensorHumidity >= 100:
		return entity.PlantMessages[3]
	case sensorHumidity >= 50 && sensorHumidity < 75:
		return entity.PlantMessages[2]
	case sensorHumidity >= 25 && sensorHumidity < 50:
		return entity.PlantMessages[1]
	case sensorHumidity < 25:
		return entity.PlantMessages[0]
	default:
		return "Impossibru"
	}
}