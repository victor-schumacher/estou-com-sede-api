package main

import (
	"github.com/victor-schumacher/estou-com-sede-api/server/http/plant"
	"log"
	"net/http"
)

func main(){
	plantManager := plant.NewPlantManager()

	http.Handle("/humidity" , plantManager.Handle())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
