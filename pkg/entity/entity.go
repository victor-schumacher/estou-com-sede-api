package entity

type Plant struct {
	SensorHumidity string
	//FinalHumidity  int `json:"finalHumidity"`
	EmailAddress string `json:"emailAddress"`
}

var PlantMessages = []string{"Estou com muita sede", "Estou com sede", "Tudo bem por aqui", "Agua demais!"}

var RandomMessages = []string{"Still waters run deep",
	"We are made for loving. If we don’t love, we will be like plants without water",
	"Como as plantas, a amizade não deve ser muito nem pouco regada.",
	"A alegria é para o corpo humano o mesmo que o sol é para as plantas.",
	"A saúde e o prazer são para o homem o que o sol e o ar são para as plantas.",
	"Meus amigos, nunca digam que há plantas más ou homens maus. O que há são maus cultivadores.",
	"Só percebemos o valor da água depois que a fonte seca.",
	"Quem quer matar a sede não procura entender a fórmula da água."}
