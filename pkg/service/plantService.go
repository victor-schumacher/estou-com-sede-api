package service

import (
	"fmt"
	"github.com/victor-schumacher/estou-com-sede-api/pkg/entity"
	"math/rand"
	"time"
)

func HumidityLevel(sensorHumidity int) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(entity.RandomMessages))

	switch {
	case sensorHumidity >= 75:
		return fmt.Sprintf("%s \n %q", entity.PlantMessages[3], entity.RandomMessages[index])
	case sensorHumidity >= 50 && sensorHumidity < 75:
		return fmt.Sprintf("%s \n %q", entity.PlantMessages[2], entity.RandomMessages[index])
	case sensorHumidity >= 25 && sensorHumidity < 50:
		return fmt.Sprintf("%s \n %q", entity.PlantMessages[1], entity.RandomMessages[index])
	case sensorHumidity < 25:
		return fmt.Sprintf("%s \n %q", entity.PlantMessages[0], entity.RandomMessages[index])
	default:
		return "Impossibru"
	}
}
