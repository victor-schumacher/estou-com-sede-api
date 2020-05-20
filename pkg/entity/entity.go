package entity

type  Plant struct {
	SensorHumidity int `json:"sensorHumidity"`
	FinalHumidity int `json:"finalHumidity"`
}

var PlantMessages = []string {"Estou morrendo de sede", "Estou com sede", "Tudo bem por aqui", "Agua demais!"}

